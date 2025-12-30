package display_tools

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"sort"
)

type KScreenDoctor struct {
}

func (ksd *KScreenDoctor) Info() string {
	return "using 'kscreen-doctor'"
}

func (ksd *KScreenDoctor) ListDisplays() ([]Display, error) {
	parsed, err := getDisplayList()
	if err != nil {
		return nil, err
	}

	displays := make([]Display, 0, len(parsed.Outputs))
	if len(parsed.Outputs) == 0 {
		return displays, nil
	}
	for i, o := range parsed.Outputs {
		displays = append(displays, Display{
			ID:        o.Name,
			Active:    o.Enabled,
			Connected: o.Connected,
			Primary:   i == 0,
		})
	}

	return displays, nil
}

func (ksd *KScreenDoctor) GetDisplay(id string) (Display, error) {
	display := Display{ID: id}
	parsed, err := getDisplayList()
	if err != nil {
		return display, err
	}

	if len(parsed.Outputs) == 0 {
		return display, fmt.Errorf("no outputs found")
	}
	for i, o := range parsed.Outputs {
		if o.Name == id {
			return Display{
				ID:        o.Name,
				Active:    o.Enabled,
				Connected: o.Connected,
				Primary:   i == 0,
			}, nil
		}
	}

	return Display{}, fmt.Errorf("display %q not found", id)
}

type kScreenDoctorJson struct {
	Outputs []struct {
		Name      string `json:"name"`
		Enabled   bool   `json:"enabled"`
		Connected bool   `json:"connected"`
		Priority  int    `json:"priority"`
	} `json:"outputs"`
}

func getDisplayList() (kScreenDoctorJson, error) {
	cmd := exec.Command("kscreen-doctor", "-j")
	out, err := cmd.Output()
	var parsed kScreenDoctorJson

	if err != nil {
		return parsed, err
	}
	if err = json.Unmarshal(out, &parsed); err != nil {
		return parsed, err
	}

	sort.Slice(parsed.Outputs, func(i, j int) bool {
		if parsed.Outputs[i].Enabled == parsed.Outputs[j].Enabled {
			return parsed.Outputs[i].Priority < parsed.Outputs[j].Priority
		}
		return parsed.Outputs[i].Enabled
	})

	return parsed, nil
}

func (ksd *KScreenDoctor) EnableDisplay(id string) error {
	cmd := exec.Command("kscreen-doctor", "output."+id+".enable")
	return cmd.Run()
}

func (ksd *KScreenDoctor) DisableDisplay(id string) error {
	cmd := exec.Command("kscreen-doctor", "output."+id+".disable")
	return cmd.Run()
}

func (ksd *KScreenDoctor) ToggleDisplay(id string) error {
	display, err := ksd.GetDisplay(id)
	if err != nil {
		return err
	}
	if display.Active {
		return ksd.DisableDisplay(id)
	} else {
		return ksd.EnableDisplay(id)
	}
}
