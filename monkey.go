package rhcos

/* This is hack a that monkey patches in run-time overrides to the defaults.
   It expects to be in the `pkg/rhcos` directory of of the installer source.
*/

import (
	"context"
	"fmt"
	"os"

	"bou.ke/monkey"
	"github.com/sirupsen/logrus"
)

// Monkey build time.
func init() {
	DefaultChannel = "ootpa"
	logrus.Infof("Overriding default channel with 'ootpa'")
	logrus.Infof(`This build allows for the following envVars to override defaults:
	RHCOS_CHANNEL: the channel to use, defaults to %s
	RHCOS_NAME: the release name, defualts to %s
	RHCOS_QCOW: use a local qcow2 image
	RHCOS_URL: the URL to fetch images, defaults to %s

	`, DefaultChannel, baseURL, buildName)

	// Replace the Channel
	newChannel, ok := os.LookupEnv("RHCOS_CHANNEL")
	if ok {
		DefaultChannel = newChannel
		logrus.Infof("Setting channel to %s", newChannel)
	}

	// Replace the URL
	newURL, ok := os.LookupEnv("RHCOS_URL")
	if ok {
		baseURL = newURL
		logrus.Infof("Setting URL to %s", newURL)
	}

	// Replace the build number
	name, ok := os.LookupEnv("RHCOS_NAME")
	if ok {
		buildName = name
		logrus.Infof("Setting build number to %s", name)
	}

	// if QCOW envVar is set, use that image.
	img, ok := os.LookupEnv("RHCOS_QCOW")
	if ok {
		monkey.Patch(QEMU, func(ctx context.Context, channel string) (string, error) {
			return fmt.Sprintf("file://%s", img), nil
		})
		logrus.Infof("Setting libvirt image to use %s", img)
	}
}
