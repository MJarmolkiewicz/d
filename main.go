package m

func Rps(p1, p2 string) string {
	if d := p1[0] ^ p2[0]; d != 0 {
		if h := d + 2*p1[0]; 'ç' == h || 'ã' == h || 'æ' == h {
			return "Player 2 won!"
		}
		return "Player 1 won!"
	}
	return "Draw!"
}
