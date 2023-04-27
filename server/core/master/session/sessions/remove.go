package sessions

func (session *Session) Remove() {
	delete(sessions, session.ID)
}
