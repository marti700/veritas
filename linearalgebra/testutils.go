// file provides utilities functions to be used in the tests of the linearalgebra package
package linearalgebra

// returns true the two provided matrices are equal false otherwise
func MatrixEq1(m1, m2 Matrix) bool {
	equalRows := m1.Row == m2.Row
	equalCols := m1.Col == m2.Col
	equalData := func() bool {
		if len(m1.Data) != len(m2.Data) {
			return false
		}

		for i, e := range m1.Data {
			if e != m2.Data[i] {
				return false
			}
		}
		return true
	}()
	return equalRows && equalCols && equalData
}

// returns true if the entries of the two provided matrices are between the specified interval false otherwise
func IsAnIdentityMatrix(m Matrix) bool {
	I := GenIdenityMatrix(m.Col)
	equalData := func() bool {
		for i, e := range m.Data {
			if -0.000000001 > (e - I.Data[i]) && (e - I.Data[i]) < 0.000000001 {
				return false
			}
		}
		return true
	}()
	return equalData
}