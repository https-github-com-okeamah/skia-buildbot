#!/usr/bin/env python
# Copyright (c) 2012 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

""" Run the Skia tests executable. """

from utils import misc
from build_step import BuildStep
import sys

class RunTests(BuildStep):
  def _Run(self, args):
    return misc.Bash([self._PathToBinary('tests')])

if '__main__' == __name__:
  sys.exit(BuildStep.Run(RunTests))