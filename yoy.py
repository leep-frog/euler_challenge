import numpy as np
from scipy.sparse import csr_matrix
from gmpy2 import mpfr as mpf, get_context

# Set the precision (e.g., 50 decimal places)
get_context().precision = 50

# Create a sparse matrix (5x5 for this example)
# Define row indices, column indices, and non-zero values
row_indices = np.array([0, 1, 2, 3, 4])
col_indices = np.array([0, 1, 2, 3, 4])
data = np.array([1.2345678901234567890123456789012345678901234567890,
                2.3456789012345678901234567890123456789012345678901,
                3.4567890123456789012345678901234567890123456789012,
                4.5678901234567890123456789012345678901234567890123,
                5.6789012345678901234567890123456789012345678901234])

# Create the sparse matrix in CSR format with float64 data type
matrix = csr_matrix((data, (row_indices, col_indices)), shape=(5, 5))

# Convert the data to gmpy2.mpf (high precision)
mpf_data = np.array([mpf(val) for val in data])

# Create a new sparse matrix where the non-zero entries are mpf objects
mpf_matrix = csr_matrix((mpf_data, (row_indices, col_indices)), shape=(5, 5))

# Print the sparse matrix in CSR format (non-zero entries will be mpf)
print("Sparse Matrix (with mpf data):")
print(mpf_matrix)

# Convert the sparse matrix to dense format to perform operations with high precision
dense_matrix = mpf_matrix.toarray()

# Now you can perform any high-precision operations using gmpy2
# Example: Add a scalar (e.g., 2.5) to each element of the dense matrix
scalar = mpf('2.5')
dense_matrix += scalar

# Print the updated dense matrix (now with high precision arithmetic)
print("\nDense Matrix After Adding Scalar:")
print(dense_matrix)
