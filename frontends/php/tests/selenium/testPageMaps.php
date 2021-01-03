<?php
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

require_once dirname(__FILE__).'/../include/class.cwebtest.php';

class testPageMaps extends CWebTest {
	public static function allMaps() {
		return DBdata('select * from sysmaps');
	}

	/**
	* @dataProvider allMaps
	*/
	public function testPageMaps_CheckLayout($map) {
		$this->zbxTestLogin('sysmaps.php');
		$this->zbxTestCheckTitle('Configuration of network maps');

		$this->zbxTestCheckHeader('Maps');

		$this->zbxTestTextPresent('Displaying');
		$this->zbxTestTextNotPresent('Displaying 0');
		$this->zbxTestTextPresent(['Name', 'Width', 'Height', 'Actions']);
		$this->zbxTestTextPresent([$map['name'], $map['width'], $map['height']]);
		$this->zbxTestTextPresent(['Delete', 'Export']);
	}

	/**
	* @dataProvider allMaps
	*/
	public function testPageMaps_SimpleEdit($map) {
		$name = $map['name'];
		$sysmapid = $map['sysmapid'];

		$sqlMap = "select * from sysmaps where name='$name' order by sysmapid";
		$oldHashMap = DBhash($sqlMap);
		$sqlElements = "select * from sysmaps_elements where sysmapid=$sysmapid order by selementid";
		$oldHashElements = DBhash($sqlElements);
		$sqlLinks = "select * from sysmaps_links where sysmapid=$sysmapid order by linkid";
		$oldHashLinks = DBhash($sqlLinks);
		$sqlLinkTriggers = "SELECT slt.* FROM sysmaps_link_triggers slt, sysmaps_links sl WHERE slt.linkid = sl.linkid AND sl.sysmapid=$sysmapid ORDER BY slt.linktriggerid";
		$oldHashLinkTriggers = DBhash($sqlLinkTriggers);

		$this->zbxTestLogin('sysmaps.php');
		$this->zbxTestCheckTitle('Configuration of network maps');
		$this->zbxTestClickLinkText($name);

		$this->zbxTestClickWait('edit');
		$this->zbxTestCheckHeader('Network maps');
		$this->zbxTestClickWait('sysmap_update');
		$this->zbxTestAcceptAlert();

		$this->zbxTestCheckTitle('Configuration of network maps');
		$this->zbxTestTextPresent($name);
		$this->zbxTestCheckHeader('Maps');

		$this->assertEquals($oldHashMap, DBhash($sqlMap), "Chuck Norris: Map update changed data in table 'sysmaps'");
		$this->assertEquals($oldHashElements, DBhash($sqlElements), "Chuck Norris: Map update changed data in table 'sysmaps_elements'");
		$this->assertEquals($oldHashLinks, DBhash($sqlLinks), "Chuck Norris: Map update changed data in table 'sysmaps_links'");
		$this->assertEquals($oldHashLinkTriggers, DBhash($sqlLinkTriggers), "Chuck Norris: Map update changed data in table 'sysmaps_link_triggers'");
	}

	/**
	* @dataProvider allMaps
	*/
	public function testPageMaps_SimpleUpdate($map) {
		$name = $map['name'];
		$sysmapid = $map['sysmapid'];

		$sqlMap = "select * from sysmaps where name='$name' order by sysmapid";
		$oldHashMap = DBhash($sqlMap);
		$sqlElements = "select * from sysmaps_elements where sysmapid=$sysmapid order by selementid";
		$oldHashElements = DBhash($sqlElements);
		$sqlLinks = "select * from sysmaps_links where sysmapid=$sysmapid order by linkid";
		$oldHashLinks = DBhash($sqlLinks);
		$sqlLinkTriggers = "select * from sysmaps_link_triggers where linkid in (select linkid from sysmaps_links where sysmapid=$sysmapid) order by linktriggerid";
		$oldHashLinkTriggers = DBhash($sqlLinkTriggers);

		$this->zbxTestLogin('sysmaps.php');
		$this->zbxTestCheckTitle('Configuration of network maps');
		$this->zbxTestClickXpathWait("//a[text()='".$name."']/../..//a[text()='Properties']");
		$this->zbxTestCheckHeader('Network maps');
		$this->zbxTestClickWait('update');
		$this->zbxTestCheckTitle('Configuration of network maps');
		$this->zbxTestWaitUntilMessageTextPresent('msg-good','Network map updated');
		$this->zbxTestTextPresent($name);
		$this->zbxTestTextPresent('Configuration of network maps');

		$this->assertEquals($oldHashMap, DBhash($sqlMap), "Chuck Norris: Map update changed data in table 'sysmaps'");
		$this->assertEquals($oldHashElements, DBhash($sqlElements), "Chuck Norris: Map update changed data in table 'sysmaps_elements'");
		$this->assertEquals($oldHashLinks, DBhash($sqlLinks), "Chuck Norris: Map update changed data in table 'sysmaps_links'");
		$this->assertEquals($oldHashLinkTriggers, DBhash($sqlLinkTriggers), "Chuck Norris: Map update changed data in table 'sysmaps_link_triggers'");
	}

	/**
	 * @dataProvider allMaps
	 * @backup sysmaps
	 */
	public function testPageMaps_MassDelete($map) {
		$sysmapid = $map['sysmapid'];

		$this->zbxTestLogin('sysmaps.php');
		$this->zbxTestCheckTitle('Configuration of network maps');
		$this->zbxTestCheckboxSelect('maps_'.$sysmapid);
		$this->zbxTestClickButton('map.massdelete');

		$this->zbxTestAcceptAlert();
		$this->zbxTestCheckTitle('Configuration of network maps');
		$this->zbxTestTextPresent('Network map deleted');

		$sql = "select * from sysmaps where sysmapid=$sysmapid";
		$this->assertEquals(0, DBcount($sql), 'Data from sysmaps table was not deleted');
		$sql = "select * from sysmaps_elements where sysmapid=$sysmapid";
		$this->assertEquals(0, DBcount($sql), 'Data from sysmaps_elements table was not deleted');
		$sql = "select * from sysmaps_links where sysmapid=$sysmapid";
		$this->assertEquals(0, DBcount($sql), 'Data from sysmaps_links table was not deleted');
		$sql = "select * from sysmaps_link_triggers where linkid in (select linkid from sysmaps_links where sysmapid=$sysmapid) order by linktriggerid";
		$this->assertEquals(0, DBcount($sql), 'Data from sysmaps_link_triggers table was not deleted');
		$sql = "select * from screens_items where resourcetype=".SCREEN_RESOURCE_MAP." and resourceid=$sysmapid;";
		$this->assertEquals(0, DBcount($sql), 'Data from screens_items table was not deleted');
	}

	public function testPageMaps_Create() {
		$this->zbxTestLogin('sysmaps.php');
		$this->zbxTestCheckTitle('Configuration of network maps');
		$this->zbxTestClickWait('form');
		$this->zbxTestTextPresent('Map');
		$this->zbxTestCheckHeader('Network maps');
		$this->zbxTestClickWait('cancel');
		$this->zbxTestCheckTitle('Configuration of network maps');
		$this->zbxTestTextPresent('Configuration of network maps');
	}

}
