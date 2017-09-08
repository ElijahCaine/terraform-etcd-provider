package etcd

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Basic Terraform Provider boilerplate
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"etcd_member": resourceMember(),
		},

		Schema: map[string]*schema.Schema{
			"endpoints": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"auto_sync_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dial_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			// TODO: v3.2.8 release
			/*
				"dial_keep_alive_time": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			*/
			// TODO: v3.2.8 release
			/*
				"dial_keep_alive_timeout": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			*/
			"tls_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tls_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tls_trusted_ca": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reject_old_cluster": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	cfg := clientv3.Config{}

	if v, ok := d.GetOk("endpoints"); ok {
		var endpoints []string
		for _, endpoint := range v.([]interface{}) {
			endpoints = append(endpoints, endpoint.(string))
		}
		cfg.Endpoints = endpoints
	}

	var err error

	if v, ok := d.GetOk("auto_sync_interval"); ok {
		if cfg.AutoSyncInterval, err = time.ParseDuration(v.(string)); err != nil {
			return nil, fmt.Errorf("Failed to parse duration: %s", err)
		}
	}

	if v, ok := d.GetOk("dial_timeout"); ok {
		if cfg.DialTimeout, err = time.ParseDuration(v.(string)); err != nil {
			return nil, fmt.Errorf("Failed to parse duration: %s", err)
		}
	}

	// TODO: v3.2.8 release
	/*
		if v, ok := d.GetOk("dial_keep_alive_time"); ok {
			if cfg.DialKeepAliveTime, err = time.ParseDuration(v.(string)); err != nil {
				return nil, fmt.Errorf("Failed to parse duration: %s", err)
			}
		}
	*/

	// TODO: v3.2.8 release
	/*
		if v, ok := d.GetOk("dial_keep_alive_timeout"); ok {
			if cfg.DialKeepAliveTimeout, err = time.ParseDuration(v.(string)); err != nil {
				return nil, fmt.Errorf("Failed to parse duration: %s", err)
			}
		}
	*/

	// Requries a cert, key, and trusted ca.
	// Is all of that entirely necessary? Not sure.
	trusted_ca, trusted_ca_ok := d.GetOk("tls_trusted_ca")
	cert, cert_ok := d.GetOk("tls_cert")
	key, key_ok := d.GetOk("tls_key")

	if trusted_ca_ok && cert_ok && key_ok {

		cfg.TLS, err = handleTLSConfig(bytes.NewBufferString(cert.(string)).Bytes(),
			bytes.NewBufferString(key.(string)).Bytes(),
			bytes.NewBufferString(trusted_ca.(string)).Bytes())

		if err != nil {
			return nil, fmt.Errorf("Failed to configure TLS: %s", err)
		}
	}

	if v, ok := d.GetOk("username"); ok {
		cfg.Username = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		cfg.Password = v.(string)
	}

	if v, ok := d.GetOk("reject_old_cluster"); ok {
		cfg.RejectOldCluster = v.(bool)
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("Failed to configure: %s", err)
	}

	return client, nil
}

func handleTLSConfig(cert_content, key_content, trusted_ca_content []byte) (*tls.Config, error) {

	// Create temp files
	cert_file, err := ioutil.TempFile("", "terraform-etcd")
	key_file, err := ioutil.TempFile("", "terraform-etcd")
	trusted_ca_file, err := ioutil.TempFile("", "terraform-etcd")

	// Defer cleanup
	defer os.Remove(cert_file.Name())
	defer os.Remove(key_file.Name())
	defer os.Remove(trusted_ca_file.Name())

	// Write to the files
	if _, err := cert_file.Write(cert_content); err != nil {
		return nil, err
	}
	if _, err := key_file.Write(key_content); err != nil {
		return nil, err
	}
	if _, err := trusted_ca_file.Write(trusted_ca_content); err != nil {
		return nil, err
	}

	// Close the files
	if err := cert_file.Close(); err != nil {
		return nil, err
	}
	if err := key_file.Close(); err != nil {
		return nil, err
	}
	if err := trusted_ca_file.Close(); err != nil {
		return nil, err
	}

	// Create the TLSInfo struct
	tlsInfo := transport.TLSInfo{
		CertFile:      cert_file.Name(),
		KeyFile:       key_file.Name(),
		TrustedCAFile: trusted_ca_file.Name(),
	}

	// Convert that to a TLSConfig
	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		return nil, err
	}

	return tlsConfig, nil
}
