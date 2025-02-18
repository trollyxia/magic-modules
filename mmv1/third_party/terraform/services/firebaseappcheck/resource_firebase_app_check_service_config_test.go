package firebaseappcheck_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"service_id":    "identitytoolkit.googleapis.com",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckFirebaseAppCheckServiceConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigUnenforcedExample(context),
			},
			{
				ResourceName:            "google_firebase_app_check_service_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id"},
			},
			{
				Config: testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigOffExample(context),
			},
			{
				ResourceName:            "google_firebase_app_check_service_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id"},
			},
			{
				Config: testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigEnforcedExample(context),
			},
			{
				ResourceName:            "google_firebase_app_check_service_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id"},
			},
		},
	})
}
