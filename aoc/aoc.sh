#!/bin/bash

_aoc_path="$(dirname -- "${BASH_SOURCE[0]}")/cmd"
_aoccmd_path="$(dirname -- "${BASH_SOURCE[0]}")/aoccmd"
aliaser aoc goleep aoc -d $_aoc_path
aliaser ac goleep aoc -d $_aoc_path d
aliaser ax goleep aoc -d $_aoc_path d -x

function at {
  pushd . > /dev/null 2>&1
  cd $_aoccmd_path
  gt
  popd > /dev/null 2>&1
}

function atv {
  pushd . > /dev/null 2>&1
  cd $_aoccmd_path
  gt -v
  popd > /dev/null 2>&1
}
