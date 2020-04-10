
func compressString(S string) string {

	s_new :=strings.Builder{}
	_len:=len(S)
	n:=0
	for i,j:=0,1;i<_len;i,j=i+1,j+1{
		n++
		if j>=_len{
			s_new.WriteRune(rune(S[i]))
			s_new.WriteString(strconv.Itoa(n))
			break
		}

		if S[i]==S[j]{
			continue
		}else{
			s_new.WriteRune(rune(S[i]))
			s_new.WriteString(strconv.Itoa(n))
			n=0
		}
	}

	s_str := s_new.String()
	if len(s_str)<_len{
		return s_str
	}
	return S
}