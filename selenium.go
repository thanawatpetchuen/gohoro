package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func NewSelenium(seleniumCfg Selenium, driverConfig ChromeDriver) (*selenium.Service, error) {
	selenium.SetDebug(false)
	opts := []selenium.ServiceOption{
		// selenium.StartFrameBuffer(),             // Start an X frame buffer for the browser to run in.
		selenium.ChromeDriver(driverConfig.Path), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),               // Output debug information to STDERR.
	}
	if seleniumCfg.Debug {
		log.Println("Debug Mode")
		selenium.SetDebug(true)
	}
	return selenium.NewSeleniumService(seleniumCfg.Path, seleniumCfg.Port, opts...)
}

func NewWebDriver(seleniumCfg Selenium, driverConfig ChromeDriver) (selenium.WebDriver, error) {
	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}

	args := []string{"--no-sandbox"}

	if driverConfig.Headless {
		args = append(args, "--headless")
	}

	if driverConfig.MuteAudio {
		args = append(args, "--mute-audio")
	}

	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: args,
	}
	caps.AddChrome(chromeCaps)
	return selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", seleniumCfg.Port))
}
