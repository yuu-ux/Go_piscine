package piscine

import "ft"

func printComma() {
	ft.PrintRune(',')
	ft.PrintRune(' ')
}

func PrintCombN(n int) {
	switch n {
	case 1:
		for i := 0; i < 10; i++ {
			ft.PrintRune(rune(i + '0'))
			if i != 9 {
				printComma()
			}
		}
	case 2:
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				ft.PrintRune(rune(i + '0'))
				ft.PrintRune(rune(j + '0'))
				if i != 8 {
					printComma()
				}
			}
		}
	case 3:
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					ft.PrintRune(rune(i + '0'))
					ft.PrintRune(rune(j + '0'))
					ft.PrintRune(rune(k + '0'))
					if i != 7 {
						printComma()
					}
				}
			}
		}
	case 4:
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {
						ft.PrintRune(rune(i + '0'))
						ft.PrintRune(rune(j + '0'))
						ft.PrintRune(rune(k + '0'))
						ft.PrintRune(rune(l + '0'))
						if i != 6 {
							printComma()
						}
					}
				}
			}
		}
	case 5:
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {
						for m := l + 1; m < 10; m++ {
							ft.PrintRune(rune(i + '0'))
							ft.PrintRune(rune(j + '0'))
							ft.PrintRune(rune(k + '0'))
							ft.PrintRune(rune(l + '0'))
							ft.PrintRune(rune(m + '0'))
							if i != 5 {
								printComma()
							}
						}
					}
				}
			}
		}
	case 6:
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {
						for m := l + 1; m < 10; m++ {
							for n := m + 1; n < 10; n++ {
								ft.PrintRune(rune(i + '0'))
								ft.PrintRune(rune(j + '0'))
								ft.PrintRune(rune(k + '0'))
								ft.PrintRune(rune(l + '0'))
								ft.PrintRune(rune(m + '0'))
								ft.PrintRune(rune(n + '0'))
								if i != 4 {
									printComma()
								}
							}
						}
					}
				}
			}
		}
	case 7:
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {
						for m := l + 1; m < 10; m++ {
							for n := m + 1; n < 10; n++ {
								for o := n + 1; o < 10; o++ {
									ft.PrintRune(rune(i + '0'))
									ft.PrintRune(rune(j + '0'))
									ft.PrintRune(rune(k + '0'))
									ft.PrintRune(rune(l + '0'))
									ft.PrintRune(rune(m + '0'))
									ft.PrintRune(rune(n + '0'))
									ft.PrintRune(rune(o + '0'))
									if i != 3 {
										printComma()
									}
								}
							}
						}
					}
				}
			}
		}
	case 8:
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {
						for m := l + 1; m < 10; m++ {
							for n := m + 1; n < 10; n++ {
								for o := n + 1; o < 10; o++ {
									for p := o + 1; p < 10; p++ {
										ft.PrintRune(rune(i + '0'))
										ft.PrintRune(rune(j + '0'))
										ft.PrintRune(rune(k + '0'))
										ft.PrintRune(rune(l + '0'))
										ft.PrintRune(rune(m + '0'))
										ft.PrintRune(rune(n + '0'))
										ft.PrintRune(rune(o + '0'))
										ft.PrintRune(rune(p + '0'))
										if i != 2 {
											printComma()
										}
									}
								}
							}
						}
					}
				}
			}
		}
	case 9:
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {
						for m := l + 1; m < 10; m++ {
							for n := m + 1; n < 10; n++ {
								for o := n + 1; o < 10; o++ {
									for p := o + 1; p < 10; p++ {
										for q := p + 1; q < 10; q++ {
											ft.PrintRune(rune(i + '0'))
											ft.PrintRune(rune(j + '0'))
											ft.PrintRune(rune(k + '0'))
											ft.PrintRune(rune(l + '0'))
											ft.PrintRune(rune(m + '0'))
											ft.PrintRune(rune(n + '0'))
											ft.PrintRune(rune(o + '0'))
											ft.PrintRune(rune(p + '0'))
											ft.PrintRune(rune(q + '0'))
											if i != 1 {
												printComma()
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
	ft.PrintRune('\n')
}
