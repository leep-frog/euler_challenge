import numpy as np
from itertools import combinations_with_replacement
from scipy.linalg import solve

def generate_states(n, k):
    """
    Generate all possible configurations of k balls on an n-node cycle.
    Each state is a sorted tuple representing ball positions.
    """
    return list(combinations_with_replacement(range(n), k))

def transition_probabilities(n, k, states):
    """
    Construct the transition probability matrix and the expected time equation matrix.
    """
    state_index = {state: i for i, state in enumerate(states)}
    num_states = len(states)
    A = np.zeros((num_states, num_states))  # Coefficients matrix
    b = np.ones(num_states)  # Constant terms (right-hand side)

    for i, state in enumerate(states):
        if len(set(state)) == 1:  # Absorbing state (all balls in one node)
            A[i, i] = 1
            b[i] = 0  # Absorbing state has expected time 0
            continue

        A[i, i] = -1  # Diagonal term (from first-step analysis)

        for ball in range(k):  # Consider moving each ball
            for direction in [-1, 1]:
                new_state = list(state)
                new_state[ball] = (new_state[ball] + direction) % n
                new_state.sort()  # Sort to maintain canonical form
                new_state = tuple(new_state)
                j = state_index[new_state]
                A[i, j] += 1 / (2 * k)

    return A, b

def expected_meeting_time(n, k, initial_config):
    """
    Compute the expected number of steps for k balls to meet in an n-node cycle.
    """
    states = generate_states(n, k)
    A, b = transition_probabilities(n, k, states)
    T_values = solve(A, b)  # Solve the linear system Ax = b

    # Get the expected time for the given initial configuration
    state_index = {state: i for i, state in enumerate(states)}
    print(T_values)
    return T_values[state_index[tuple(sorted(initial_config))]]

# Example usage
n = 10  # Number of nodes in the cycle
k = 8  # Number of balls
initial_config = (0, 0, 0, 0, 0, 0, 1, 2)  # Initial positions of the balls

print("START")
expected_time = expected_meeting_time(n, k, initial_config)
print(f"Expected steps until all balls meet: {expected_time:.4f}")
