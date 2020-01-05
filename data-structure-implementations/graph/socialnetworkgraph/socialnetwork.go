package socialnetworkgraph

import (
	"sync"
)

type FriendNode struct{
	profileId int
	name string
	nextFriend *FriendNode
}

type GraphKeyNode struct {
	profileId int
	name string
	nextKey *GraphKeyNode
	nodeMutex sync.RWMutex
	friendList *FriendNode
}

type SocialNetworkGraph struct {
	graphMutex sync.RWMutex
	head *GraphKeyNode
	size int
}

type SocialNetworkGraphInterface interface {
	CreateNewSocialNetworkGraph()
	AddNewPerson(profileId int, name string) *GraphKeyNode
	FindGraphSize()
	AddFriend(person *GraphKeyNode , friend *FriendNode)
	FindFriendList(person *GraphKeyNode)
	FindMutualFriends(person1 *GraphKeyNode, person2 *GraphKeyNode)
}