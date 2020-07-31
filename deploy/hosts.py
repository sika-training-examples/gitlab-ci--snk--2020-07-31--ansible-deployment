#!/usr/bin/env python3

import json
import os
import sys

host = os.environ.get("DEPLOY_HOST")
if not host:
    sys.stderr.write("ERROR: Missing value for env variable DEPLOY_HOST\n")
    sys.stderr.flush()
    quit(1)

print(json.dumps({"all": {"hosts": [host],}}))
