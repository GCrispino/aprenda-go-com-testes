package poker

import "testing"

// StubPlayerStore implements PlayerStore for testing purposes
type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []Player
}

// ObterPontuacaoDeJogador returns a score from Scores
func (s *StubPlayerStore) ObterPontuacaoDeJogador(name string) int {
	score := s.Scores[name]
	return score
}

// RecordWin will record a win to WinCalls
func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

// ObterLiga returns League
func (s *StubPlayerStore) ObterLiga() League {
	return s.League
}

// AssertPlayerWin allows you to spy on the armazenamento's calls to RecordWin
func AssertPlayerWin(t *testing.T, armazenamento *StubPlayerStore, winner string) {
	t.Helper()

	if len(armazenamento.WinCalls) != 1 {
		t.Fatalf("recebi %d chamadas de RecordWin esperava %d", len(armazenamento.WinCalls), 1)
	}

	if armazenamento.WinCalls[0] != winner {
		t.Errorf("não armazenou o vencedor correto recebi '%s' esperava '%s'", armazenamento.WinCalls[0], winner)
	}
}
