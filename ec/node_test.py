import unittest
from parameterized import parameterized
import node

class TestChallenges(unittest.TestCase):

  @parameterized.expand([
      (68, [3], '432621513'),
      (68, [5], '6531031914842725'),
  ])
  def test_nodes(self, num, args, want):
    self.assertEqual(node.nodes().run_challenge(num, *args), want)
