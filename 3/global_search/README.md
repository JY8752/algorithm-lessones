# ビット演算による全探索

- 求めたい部分集合は二進数表現に対応づけることができる
- 整数列の要素数がNであれば、部分集合の全範囲は2^N通りであり、これは**1 << N**と表現できる
- 二進数で対応づけた部分集合とi番目の要素aiの**AND和**を取り、i番目の要素aiと一致すればそれは部分集合に含まれる
- こうして、部分集合に含まれる要素の和を求めてWと一致していれば、そのときの部分集合が答え