/*
** Zabbix
** Copyright (C) 2001-2021 Zabbix SIA
**
** This program is free software; you can redistribute it and/or modify
** it under the terms of the GNU General Public License as published by
** the Free Software Foundation; either version 2 of the License, or
** (at your option) any later version.
**
** This program is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
** GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License
** along with this program; if not, write to the Free Software
** Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
**/

package memcached

import (
	"reflect"
	"testing"
)

func TestURI_Scheme(t *testing.T) {
	type fields struct {
		scheme   string
		host     string
		port     string
		socket   string
		password string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Should return scheme from URI structure",
			fields{scheme: "tcp"},
			"tcp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URI{
				scheme:   tt.fields.scheme,
				host:     tt.fields.host,
				port:     tt.fields.port,
				socket:   tt.fields.socket,
				password: tt.fields.password,
			}
			if got := u.Scheme(); got != tt.want {
				t.Errorf("URI.Scheme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURI_Addr(t *testing.T) {
	type fields struct {
		scheme   string
		host     string
		port     string
		socket   string
		password string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Should return host:port from URI structure",
			fields{host: "127.0.0.1", port: "11211"},
			"127.0.0.1:11211",
		},
		{
			"Should return socket from URI structure",
			fields{socket: "unix:///var/run/memcached.sock"},
			"unix:///var/run/memcached.sock",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URI{
				scheme:   tt.fields.scheme,
				host:     tt.fields.host,
				port:     tt.fields.port,
				socket:   tt.fields.socket,
				password: tt.fields.password,
			}
			if got := u.Addr(); got != tt.want {
				t.Errorf("URI.Addr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURI_User(t *testing.T) {
	type fields struct {
		scheme   string
		host     string
		port     string
		socket   string
		user     string
		password string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Should return username from URI structure",
			fields{user: "zabbix"},
			"zabbix",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URI{
				scheme:   tt.fields.scheme,
				host:     tt.fields.host,
				port:     tt.fields.port,
				socket:   tt.fields.socket,
				user:     tt.fields.user,
				password: tt.fields.password,
			}
			if got := u.User(); got != tt.want {
				t.Errorf("URI.User() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURI_Password(t *testing.T) {
	type fields struct {
		scheme   string
		host     string
		port     string
		socket   string
		password string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Should return password from URI structure",
			fields{password: "sEcReT"},
			"sEcReT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URI{
				scheme:   tt.fields.scheme,
				host:     tt.fields.host,
				port:     tt.fields.port,
				socket:   tt.fields.socket,
				password: tt.fields.password,
			}
			if got := u.Password(); got != tt.want {
				t.Errorf("URI.Password() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURI_URI(t *testing.T) {
	type fields struct {
		scheme   string
		host     string
		port     string
		socket   string
		user     string
		password string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Should return URI with creds",
			fields{scheme: "tcp", host: "127.0.0.1", port: "11211", user: "zabbix", password: "sEcReT"},
			"tcp://zabbix:sEcReT@127.0.0.1:11211",
		},
		{
			"Should return URI without creds",
			fields{scheme: "tcp", host: "127.0.0.1", port: "11211"},
			"tcp://127.0.0.1:11211",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URI{
				scheme:   tt.fields.scheme,
				host:     tt.fields.host,
				port:     tt.fields.port,
				socket:   tt.fields.socket,
				user:     tt.fields.user,
				password: tt.fields.password,
			}
			if got := u.URI(); got != tt.want {
				t.Errorf("URI.URI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newURIWithCreds(t *testing.T) {
	type args struct {
		uri      string
		user     string
		password string
	}

	tests := []struct {
		name    string
		args    args
		wantU   *URI
		wantErr bool
	}{
		{
			"Should return URI structure",
			args{"tcp://192.168.1.1:11211", "zabbix", "sEcReT"},
			&URI{scheme: "tcp", host: "192.168.1.1", port: "11211", user: "zabbix", password: "sEcReT"},
			false,
		},
		{
			"Should return error if failed to parse URI",
			args{"://", "zabbix", "sEcReT"},
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU, err := newURIWithCreds(tt.args.uri, tt.args.user, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("newURIWithPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("newURIWithPassword() = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}

func Test_parseURI(t *testing.T) {
	type args struct {
		uri string
	}

	tests := []struct {
		name    string
		args    args
		wantU   *URI
		wantErr bool
	}{
		{
			"Parse URI with tcp scheme",
			args{"tcp://localhost:11211"},
			&URI{scheme: "tcp", host: "localhost", port: "11211"},
			false,
		},
		{
			"Parse URI without scheme",
			args{"localhost:11211"},
			nil,
			true,
		},
		{
			"Parse URI with ipv6 address. Test 1",
			args{"tcp://[fe80::1ce7:d24a:97f0:3d83%25en0]:11211"},
			&URI{scheme: "tcp", host: "fe80::1ce7:d24a:97f0:3d83%en0", port: "11211"},
			false,
		},
		{
			"Parse URI with ipv6 address. Test 2",
			args{"tcp://[fe80::1ce7:d24a:97f0:3d83%en0]:11211"},
			&URI{scheme: "tcp", host: "fe80::1ce7:d24a:97f0:3d83%en0", port: "11211"},
			false,
		},
		{
			"Parse URI with ipv6 address. Test 3",
			args{"tcp://[fe80::1%25lo0]:11211"},
			&URI{scheme: "tcp", host: "fe80::1%lo0", port: "11211"},
			false,
		},
		{
			"Parse URI with ipv6 address. Test 4",
			args{"tcp://[::1]"},
			&URI{scheme: "tcp", host: "::1", port: "11211"},
			false,
		},
		{
			"Parse URI with ipv6 address. Test 5",
			args{"tcp://fe80::1:11211"},
			&URI{scheme: "tcp", host: "fe80::1", port: "11211"},
			false,
		},
		{
			"Parse URI with ipv6 address. Test 6",
			args{"tcp://::1:11289"},
			&URI{scheme: "tcp", host: "::1", port: "11289"},
			false,
		},
		{
			"Should fail when URI without host is passed",
			args{"tcp://:11211"},
			nil,
			true,
		},
		{
			"Should fail if port is greater than 65535",
			args{"tcp://localhost:65536"},
			nil,
			true,
		},
		{
			"Should fail if port is not integer",
			args{"tcp://:foo"},
			nil,
			true,
		},
		{
			"Parse URI without port",
			args{"tcp://localhost"},
			&URI{scheme: "tcp", host: "localhost", port: "11211"},
			false,
		},
		{
			"Parse URI with unix scheme",
			args{"unix:///var/run/memcached.sock"},
			&URI{scheme: "unix", socket: "/var/run/memcached.sock"},
			false,
		},
		{
			"Should fail if an unknown scheme is passed",
			args{"foo://bar"},
			nil,
			true,
		},
		{
			"Should fail if a wrong format URI is passed",
			args{"!@#$%^&*()"},
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU, err := parseURI(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("parseURI() = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}

func Test_validateURI(t *testing.T) {
	type args struct {
		uri string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Should pass when a valid URI is given",
			args{"tcp://localhost:11211"},
			false,
		},
		{
			"Should fail when a malformed URI is given",
			args{"tcp:/localhost:11211"},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateURI(tt.args.uri); (err != nil) != tt.wantErr {
				t.Errorf("validateURI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isLooksLikeURI(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Should return true if it is URI with tcp scheme",
			args{"tcp://localhost:11211"},
			true,
		},
		{
			"Should return true if it is URI with unix scheme",
			args{"unix:/var/run/memcached.sock"},
			true,
		},
		{
			"Should return false if it is not URI",
			args{"testSession"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLooksLikeURI(tt.args.s); got != tt.want {
				t.Errorf("isLooksLikeURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
