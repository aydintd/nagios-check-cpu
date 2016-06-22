# nagios-check-cpu
check-cpu plugin for Nagios implemented in Go

## Installation

	$ go get github.com/aydintd/nagios-check-cpu
	$ cd $GOPATH/src/github.com/aydintd/nagios-check-cpu ; go install
	$ nagios-check-cpu -w 20 -c 25                                                                                                                                                                                              
	CPU: CRITICAL - TotalAvg: %89.447236|avg=%89.447236;;;; user=2512842;;;; nice=6827;;;; system=422840;;;; idle=43293876;;;; 

## LICENSE

	nagios-cpu-check - Nagios CPU check plugin
    Copyright (C) 2016  Aydin Doyak <aydintd@gmail.com>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
