"use strict";

/***
 *  Define the app and inject any modules we wish to
 *  refer to.
***/
var app = angular.module("MeetingMinutes", []);

/******************************************************************************\
Function:
    AppController

Dependencies:
    ...

Description:
    Main application controller
\******************************************************************************/
app.controller("AppController", ["$scope", "$location", "$timeout",

    function($scope, $location, $timeout) {

        // Define some dummy meetings so we can mock up the UI
        $scope.meetings = [
            { name: "test 01", date: new Date() },
            { name: "test 02", date: new Date() }
        ]

        $scope.gotoMeeting = function(meeting) {
            console.log("GOTO Meeting: " + meeting.name);
        }

    }]

);

/*****************************************************************************\
    Service to hit the server
\*****************************************************************************/
app.service("HttpService", ["$http",

    function($http) {
        return {
        };
    }]

);

