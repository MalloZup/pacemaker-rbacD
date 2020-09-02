package main

import (
	"os/exec"

	log "github.com/sirupsen/logrus"

	"time"
)

// this function in future should read a yaml file where we define:
// role name attribte xpath
// something like this role monitor read xpath:"/cib"
func initializeRoles() error {
	// this command is not idempotent, if the role exists it fails
	out, err := exec.Command("/usr/sbin/crm", "configure", "role", "haclient-readonly", "read", "xpath:\"/cib\"").CombinedOutput()
	if err != nil {
		log.Errorf("could not configure haclient-readonly via crm. %s \n error %s", out, err)
		return err
	}
	return nil
}

/// this function retrive all users which belong to group we do RBAC see initialize roles, and we add the needed acls
func assignFromGroupToRole(group string) error {

	// retrieve all user belonging to group
	user := "foo"
	// this should run in a for
	out, err := exec.Command("/usr/sbin/crm", "configure", "acl_target", user, group).CombinedOutput()
	if err != nil {
		log.Errorf("could not add user to haclient-readonly role. %s \n error %s", out, err)
		return err
	}
	return nil
}

func main() {

	for {

		err := initializeRoles()
		if err != nil {
			log.Warn(err)
		}

		err = assignFromGroupToRole("haclient-readonly")
		if err != nil {
			log.Warn(err)
		}

		log.Infoln("sleeping for 5 seconds")
		time.Sleep(5 * time.Second)
	}
}
