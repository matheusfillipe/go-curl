#!/usr/bin/env python3
# -*- coding: utf-8 -*-


""" codegen.py reads curl project's curl.h, outputting go-curl's compat.h

The purpose of compatggen.py is to generate compat.h defines which are needed
by go-curl build process. Most users of co-curl can stop reading here,
because a usable compat.h is already included in the go-curl source.

compatgen.py should be run from the go-curl source root, where it will
attempt to locate 'curl.h' in the locations defined by `target_dirs`.


CURL_GIT_PATH (if defined) must point to the location of your curl source
repository. For example you might check out curl from
https://github.com/curl/curl and have that saved to your local
directory `/Users/example-home/git/curl`

Usage:

CURL_GIT_PATH=/path-to-git-repos/curl python ./misc/compatgen.py

Example:
CURL_GIT_PATH=/Users/example-user/Projects/c/curl ./misc/compatgen.py

File Input:
(curl project) include/curl/header.h

File Output:
compat.h

Todo:
* Further code review ("help wanted")
* More docstrings/help. Error checking. Cleanup redefined variable scopes.
"""

import os
import re

# CURL_GIT_PATH is the location you git checked out the curl project.
# You will need to supply this variable and value when invoking this script.
CURL_GIT_PATH = os.environ.get("CURL_GIT_PATH", "./curl")

target_dirs = [
    "{}/include/curl".format(CURL_GIT_PATH),
    "/usr/local/include",
    "libdir/gcc/target/version/include" "/usr/target/include",
    "/usr/include",
]


def get_curl_path() -> str:
    for d in target_dirs:
        for root, dirs, files in os.walk(d):
            if "curl.h" in files:
                return os.path.join(root, "curl.h")
    raise Exception("Not found")


def version_symbol(ver: str) -> tuple[list[str], list[str], list[str], list[str], list[str]]:
    """Returns lists of symbol info, when given argument of curl's Git tag."""

    # force switch (discard changes) and checkout each of the curl branches
    checkout_cmd = 'cd "{}" && git status --porcelain && git -c advice.detachedHead=false checkout -f "{}"'.format(
        CURL_GIT_PATH, ver
    )

    os.system(checkout_cmd)
    opts = []
    codes = []
    infos = []
    vers = []
    auths = []
    init_pattern = re.compile(r"CINIT\((.*?),\s*(LONG|OBJECTPOINT|FUNCTIONPOINT|STRINGPOINT|OFF_T),\s*(\d+)\)")
    error_pattern = re.compile(r"^\s+(CURLE_[A-Z_0-9]+),")
    info_pattern = re.compile(r"^\s+(CURLINFO_[A-Z_0-9]+)\s+=")
    with open(os.path.join(CURL_GIT_PATH, "include", "curl", "curl.h")) as f:
        for line in f:
            match = init_pattern.findall(line)
            if match:
                opts.append("CURLOPT_" + match[0][0])
            if line.startswith("#define CURLOPT_"):
                o = line.split()
                opts.append(o[1])

            if line.startswith("#define CURLAUTH_"):
                a = line.split()
                auths.append(a[1])

            match = error_pattern.findall(line)
            if match:
                codes.append(match[0])

            if line.startswith("#define CURLE_"):
                c = line.split()
                codes.append(c[1])

            match = info_pattern.findall(line)
            if match:
                infos.append(match[0])

            if line.startswith("#define CURLINFO_"):
                i = line.split()
                if "0x" not in i[2]:  # :(
                    infos.append(i[1])

            if line.startswith("#define CURL_VERSION_"):
                i = line.split()
                vers.append(i[1])

    return opts, codes, infos, vers, auths


def extract_version(tag_str: str) -> dict:
    """Converts a curl git tag (curl-8_8_0) into a dict, ex.: {'major': 8, 'minor': 8, 'patch': 0, 'version': 'curl-8_8_0'}"""
    fields = re.search(r"curl-([0-9]+)_([0-9]+)_([0-9]+)", tag_str)
    if fields is None:
        raise ValueError("Invalid tag: {}".format(tag_str))
    version = {
        "major": int(fields.group(1)),
        "minor": int(fields.group(2)),
        "patch": int(fields.group(3)),
        "version": tag_str,
    }
    return version


## valid versions that are compatible are 7_16_XXX or higher
def is_valid_version(version: dict) -> bool:
    """Returns false when curl version (branch) is less than 7_16_XXX."""
    return version["major"] >= 8 or (version["major"] == 7 and version["minor"] >= 16)


# fetches a list of tags from the curl repo, matching only format "curl-X_Y_Z"
get_tags_raw_cmd = "cd {} && git tag | grep -E '^curl-[0-9]+_[0-9]+_[0-9]+$'".format(CURL_GIT_PATH)
tags_raw = os.popen(get_tags_raw_cmd).read().split("\n")[:-1]
tags = map(extract_version, tags_raw)
tags = filter(is_valid_version, tags)
versions = sorted(tags, key=lambda v: [v["major"], v["minor"], v["patch"]], reverse=True)
last = version_symbol("master")

template = """
/* generated by compatgen.py */
#include<curl/curl.h>


"""

result = [template]
result_tail = ["/* generated ends */\n"]


if __name__ == "__main__":
    for ver in versions:
        major = ver["major"]
        minor = ver["minor"]
        patch = ver["patch"]
        opts, codes, infos, vers, auths = curr = version_symbol(ver["version"])

        for o in last[0]:
            if o not in opts:
                result.append("#define {} 0".format(o))  # 0 for nil option
        for c in last[1]:
            if c not in codes:
                result.append("#define {} -1".format(c))  # -1 for error
        for i in last[2]:
            if i not in infos:
                result.append("#define {} 0".format(i))  # 0 for nil
        for v in last[3]:
            if v not in vers:
                result.append("#define {} 0".format(v))  # 0 for nil
        for a in last[4]:
            if a not in auths:
                result.append("#define {} 0".format(a))  # 0 for nil

        result.append(
            f"#if LIBCURL_VERSION_MAJOR < {major} || (LIBCURL_VERSION_MAJOR == {major} && LIBCURL_VERSION_MINOR < {minor}) || (LIBCURL_VERSION_MAJOR == {major} && LIBCURL_VERSION_MINOR == {minor} && LIBCURL_VERSION_PATCH < {patch})"
        )

        result_tail.insert(0, "#endif /* {}.{}.{} */".format(major, minor, patch))

        last = curr

result.append("#error your version is TOOOOOOOO low")

result.extend(result_tail)

with open("./compat.h", "w", encoding="utf-8") as fp:
    fp.write("\n".join(result))
