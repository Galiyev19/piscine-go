package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	if n >= 1 && n <= 9 {
		if n == 1 {
			for i := 48; i < 58; i++ {
				z01.PrintRune(rune(i))
				if rune(i) != rune(57) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		} else if n == 2 {
			for i := 48; i < 58; i++ {
				for j := i + 1; j < 58; j++ {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					if rune(i) != rune(56) {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
			}
		} else if n == 3 {
			for i := 48; i < 58; i++ {
				for j := i + 1; j < 58; j++ {
					for k := j + 1; k < 58; k++ {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(k))
						if rune(i) != rune(55) {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}
				}
			}
		} else if n == 4 {
			for i := 48; i < 58; i++ {
				for j := i + 1; j < 58; j++ {
					for k := j + 1; k < 58; k++ {
						for l := k + 1; l < 58; l++ {
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(rune(k))
							z01.PrintRune(rune(l))
							if rune(i) != rune(54) {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		} else if n == 5 {
			for i := 48; i < 58; i++ {
				for j := i + 1; j < 58; j++ {
					for k := j + 1; k < 58; k++ {
						for l := k + 1; l < 58; l++ {
							for m := l + 1; m < 58; m++ {
								z01.PrintRune(rune(i))
								z01.PrintRune(rune(j))
								z01.PrintRune(rune(k))
								z01.PrintRune(rune(l))
								z01.PrintRune(rune(m))
								if rune(i) != rune(53) {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}

							}
						}
					}
				}
			}
		} else if n == 6 {
			for i := 48; i < 58; i++ {
				for j := i + 1; j < 58; j++ {
					for k := j + 1; k < 58; k++ {
						for l := k + 1; l < 58; l++ {
							for m := l + 1; m < 58; m++ {
								for n := m + 1; n < 58; n++ {
									z01.PrintRune(rune(i))
									z01.PrintRune(rune(j))
									z01.PrintRune(rune(k))
									z01.PrintRune(rune(l))
									z01.PrintRune(rune(m))
									z01.PrintRune(rune(n))
									if rune(i) != rune(52) {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}

								}
							}
						}
					}
				}
			}
		} else if n == 7 {
			for i := 48; i < 58; i++ {
				for j := i + 1; j < 58; j++ {
					for k := j + 1; k < 58; k++ {
						for l := k + 1; l < 58; l++ {
							for m := l + 1; m < 58; m++ {
								for n := m + 1; n < 58; n++ {
									for o := n + 1; o < 58; o++ {
										z01.PrintRune(rune(i))
										z01.PrintRune(rune(j))
										z01.PrintRune(rune(k))
										z01.PrintRune(rune(l))
										z01.PrintRune(rune(m))
										z01.PrintRune(rune(n))
										z01.PrintRune(rune(o))
										if rune(i) != rune(51) {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}

									}
								}
							}
						}
					}
				}
			}
		} else if n == 8 {
			for i := 48; i < 58; i++ {
				for j := i + 1; j < 58; j++ {
					for k := j + 1; k < 58; k++ {
						for l := k + 1; l < 58; l++ {
							for m := l + 1; m < 58; m++ {
								for n := m + 1; n < 58; n++ {
									for o := n + 1; o < 58; o++ {
										for p := o + 1; p < 58; p++ {
											z01.PrintRune(rune(i))
											z01.PrintRune(rune(j))
											z01.PrintRune(rune(k))
											z01.PrintRune(rune(l))
											z01.PrintRune(rune(m))
											z01.PrintRune(rune(n))
											z01.PrintRune(rune(o))
											z01.PrintRune(rune(p))
											if rune(i) != rune(50) {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}

										}
									}
								}
							}
						}
					}
				}
			}
		} else if n == 9 {
			for i := 48; i < 58; i++ {
				for j := i + 1; j < 58; j++ {
					for k := j + 1; k < 58; k++ {
						for l := k + 1; l < 58; l++ {
							for m := l + 1; m < 58; m++ {
								for n := m + 1; n < 58; n++ {
									for o := n + 1; o < 58; o++ {
										for p := o + 1; p < 58; p++ {
											for u := p + 1; u < 58; u++ {
												z01.PrintRune(rune(i))
												z01.PrintRune(rune(j))
												z01.PrintRune(rune(k))
												z01.PrintRune(rune(l))
												z01.PrintRune(rune(m))
												z01.PrintRune(rune(n))
												z01.PrintRune(rune(o))
												z01.PrintRune(rune(p))
												z01.PrintRune(rune(u))
												if rune(i) != rune(49) {
													z01.PrintRune(',')
													z01.PrintRune(' ')
												}

											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
