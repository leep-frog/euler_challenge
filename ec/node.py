

import string


class Challenges(object):

  def __init__(self, challenges={}):
    self.challenges = challenges

  def run_challenge(self, challenge, *args):
    return self.challenges[challenge](*args)


def nodes():
  return Challenges({
    68: p68,
  })


def p68(n):
  s = [2*n-k for k in range(2*n)]

  cycles = []
  for a in s:
    sb = [k for k in s if k != a]
    for b in sb:
      sc = [k for k in sb if k != b]
      for c in sc:
        sum = a + b + c
        new_cycles = next_cycle([k for k in sc if k != c], sum, b, c, [[a, b, c]])
        for cycle in new_cycles:
          smallest = min([p[0] for p in cycle])
          start_idx = [part[0] for part in cycle].index(smallest)
          cycle = cycle[start_idx:] + cycle[:start_idx]
          cycles.append(cycle)
  
  # sort by smallest outer digit
  cycles = sorted(cycles, key=lambda cycle: min([p[0] for p in cycle]))
  cycles = [''.join([str(num) for parts in cycle for num in parts]) for cycle in cycles]
  cycles.sort()
  cycles = [cycle for cycle in cycles if len(cycle) <= 16]
  return cycles[-1]

def next_cycle(remaining, sum, last, prev, cur):
  if len(remaining) == 1:
    if last + prev + remaining[0] == sum:
      cur.append([remaining[0], prev, last])
      return [cur]
    return None

  rs = []
  for b in remaining:    
    sc = [k for k in remaining if k != b]
    c = sum - b - prev
    try:
      sc.index(c)
    except:
      continue
    cp = cur.copy()
    cp.append([b, prev, c])
    has_cycle = next_cycle([k for k in sc if k != c], sum, last, c, cp)
    if has_cycle is None:
      continue
    rs += has_cycle
  return rs

def primes(n):
  primes = [2, 3]
  while primes[-1] < 1_000_000:
    