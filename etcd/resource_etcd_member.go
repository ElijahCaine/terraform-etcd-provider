package etcd

import (
	"fmt"
	"strconv"

	"github.com/coreos/etcd/clientv3"
	"github.com/hashicorp/terraform/helper/schema"
	"golang.org/x/net/context"
)

func resourceMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceMemberCreate,
		Read:   resourceMemberRead,
		Update: resourceMemberUpdate,
		Delete: resourceMemberDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"peer_urls": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func generatePeerUrls(d *schema.ResourceData) []string {

	var peerUrls []string

	if v, ok := d.GetOk("peer_urls"); ok {
		for _, url := range v.([]interface{}) {
			peerUrls = append(peerUrls, url.(string))
		}
	}

	return peerUrls
}

func resourceMemberCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*clientv3.Client)

	resp, err := client.MemberAdd(context.Background(), generatePeerUrls(d))
	if err != nil {
		return err
	}

	memberIdStr := strconv.FormatUint(resp.Member.ID, 10)
	d.SetId(memberIdStr)

	return nil
}

func resourceMemberRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*clientv3.Client)

	resp, err := client.MemberList(context.Background())
	if err != nil {
		return err
	}

	// TODO: Add new member to the cluster config
	for _, member := range resp.Members {
		if fmt.Sprintf("%d", member.ID) == d.Id() {
			return nil
		}
	}
	d.SetId("")

	return nil
}

func resourceMemberUpdate(d *schema.ResourceData, m interface{}) error {

	d.Partial(true)

	if d.HasChange("peer_urls") {

		memberIdInt, err := strconv.ParseUint(d.Id(), 0, 64)
		if err != nil {
			return err
		}

		client := m.(*clientv3.Client)

		_, err = client.MemberUpdate(context.Background(), memberIdInt, generatePeerUrls(d))
	}
	return nil
}

func resourceMemberDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*clientv3.Client)

	memberId, err := strconv.ParseUint(d.Id(), 0, 64)
	if err != nil {
		return err
	}

	_, err = client.MemberRemove(context.Background(), memberId)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
