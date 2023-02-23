package password

// TODO: Write unit test for 'Hash', 'Check' methods and for 'SHA256Hasher' function

type Manager struct {
	hasher Hasher
	key    string
}

func NewManager(hasher Hasher, key string) *Manager {
	return &Manager{
		hasher: hasher,
		key:    key,
	}
}

func (m *Manager) Hash(password string) string {
	return m.hasher(password, m.key)
}

func (m *Manager) Check(password, hashedPassword string) bool {
	return m.hasher(password, m.key) == hashedPassword
}
