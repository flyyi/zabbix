// +build !windows

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

package zabbixsync

func getMetrics() []string {
	return []string{
		"net.dns", "Checks if DNS service is up.",
		"net.dns.record", "Performs DNS query.",
		"proc.mem", "Memory used by process in bytes.",
		"proc.num", "The number of processes.",
		"system.hw.chassis", "Chassis information.",
		"system.hw.devices", "Listing of PCI or USB devices.",
		"system.sw.packages", "Listing of installed packages.",
		"vfs.dir.count", "Directory entry count.",
		"vfs.dir.size", "Directory size (in bytes).",
		"vm.memory.size", "Memory size in bytes or in percentage from total.",
	}
}
