// Package core contains code related to this amazing custom blockchain including miners and a PoW algorithm
package core

/**
 * The MIT License (MIT)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
import "testing"

func TestHelloWorld_Verify(t *testing.T) {
	pow := ProofOfWork{}

	work := "These pretzels are making me thirsty"
	solution := "0086b692327e3a60b4884ed0f91853ccb16bd38d197df69da6bfc660820a4c0aeeac5d35a911ac2e14ce54b402cce0ae7c00c3a01c976ea54cd18902f0a6855e"

	proof := pow.Solve(work, 2)

	t.Logf("Solution: %s", proof.Solution)
	t.Logf("Variation: %s", proof.Variation)

	verified := pow.Verify(work, 2, proof)

	if proof.Solution != solution {
		t.Errorf("The solution from the proof is different from the expected solution -- %s vs %s --", proof.Solution, solution)
	}

	if verified != true {
		t.Errorf("Solution -- %s -- is incorrect. It should be -- %s --", proof.Solution, solution)
	}
}
