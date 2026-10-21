package service

import (
	"sync"

	"github.com/noirbyss/botkenc/internal/domain"
)

type navigator struct {
	mu          sync.Mutex
	usersStacks map[int64][]domain.ScreenRef
}

func newNavigator() *navigator {
	return &navigator{
		usersStacks: make(map[int64][]domain.ScreenRef),
	}
}

func (n *navigator) push(telegramUserID int64, screen domain.ScreenRef) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if screen.ID == domain.MainScreen {
		n.usersStacks[telegramUserID] = []domain.ScreenRef{screen}
		return
	}

	userStack := n.usersStacks[telegramUserID]
	n.usersStacks[telegramUserID] = append(userStack, screen)
}

func (n *navigator) back(telegramUserID int64) domain.ScreenRef {
	n.mu.Lock()
	defer n.mu.Unlock()

	userStack := n.usersStacks[telegramUserID]
	userStack = userStack[:len(userStack)-1]

	n.usersStacks[telegramUserID] = userStack

	return userStack[len(userStack)-1]
}
