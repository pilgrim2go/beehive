/*
 *    Copyright (C) 2014 Daniel 'grindhold' Brendle
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Daniel 'grindhold' Brendle <grindhold@skarphed.org>
 */

package nagiosbee

import (
	"github.com/muesli/beehive/modules"
)

type NagiosBeeFactory struct{}

func (factory *NagiosBeeFactory) New(name, description string, options modules.BeeOptions) modules.ModuleInterface {
	bee := NagiosBee{
		name:        name,
		namespace:   factory.Name(),
		url:         options.GetValue("url").(string),
		user:        options.GetValue("user").(string),
		password:    options.GetValue("password").(string),
		description: description,
	}
	return &bee
}

func (factory *NagiosBeeFactory) Name() string {
	return "nagiosbee"
}

func (factory *NagiosBeeFactory) Description() string {
	return "A bee that fetches status changes from nagios-monitors."
}

func (factory *NagiosBeeFactory) Options() []modules.BeeOptionDescriptor {
	opts := []modules.BeeOptionDescriptor{
		modules.BeeOptionDescriptor{
			Name:        "url",
			Description: "URL to the statusJson.php-script typically http://domain.com/nagios3/statusJson.php",
			Type:        "string",
		},
		modules.BeeOptionDescriptor{
			Name:        "user",
			Description: "The username of the nagios-user",
			Type:        "string",
		},
		modules.BeeOptionDescriptor{
			Name:        "password",
			Description: "Password of the nagios-user's account",
			Type:        "string",
		},
	}
	return opts
}

func (factory *NagiosBeeFactory) Events() []modules.EventDescriptor {
	events := []modules.EventDescriptor{
		modules.EventDescriptor{
			Namespace:   factory.Name(),
			Name:        "statuschange",
			Description: "The status of a Service has changed",
			Options: []modules.PlaceholderDescriptor{
				modules.PlaceholderDescriptor{
					Name:        "server",
					Description: "Name of the system the changed server resides on",
					Type:        "string",
				},
				modules.PlaceholderDescriptor{
					Name:        "service",
					Description: "Name of the service that has changed",
					Type:        "string",
				},
				modules.PlaceholderDescriptor{
					Name:        "message",
					Description: "Message that the NRPE-service returned",
					Type:        "string",
				},
				modules.PlaceholderDescriptor{
					Name:        "status",
					Description: "New status number",
					Type:        "int",
				},
			},
		},
	}
	return events
}

func (factory *NagiosBeeFactory) Actions() []modules.ActionDescriptor {
	actions := []modules.ActionDescriptor{}
	return actions
}

func init() {
	f := NagiosBeeFactory{}
	modules.RegisterFactory(&f)
}
