"use strict";

/***
 *  Define the app and inject any modules we wish to
 *  refer to.
***/
var app = angular.module("MeetingMinutes", [
    "ngRoute"
]);

/***
 * Define the routes which the application should support
***/
app.config(["$routeProvider",

    function($routeProvider) {
        $routeProvider.
            when("/meetings", {
                templateUrl: "partials/meeting-list.html",
                controller:  "MeetingListController"
            }).when("/meetings/:meetingId", {
                templateUrl: "partials/meeting.html",
                controller:  "MeetingController"
            }).otherwise({ redirectTo: "/meetings" });
    }]

);

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

/******************************************************************************\
Function:
    MeetingListController 

Dependencies:
    ...

Description:
    Application controller responsible for querying and displaying meetings
\******************************************************************************/
app.controller("MeetingListController", ["$scope", "$location", "$timeout",

    function($scope, $location, $timeout) {
        console.log("Meeting List Controller initialized");
    }]

);


/******************************************************************************\
Function:
    MeetingController 

Dependencies:
    ...

Description:
    Application controller responsible for rendering a single meeting 
\******************************************************************************/
app.controller("MeetingController", ["$scope", "$location", "$timeout",

    function($scope, $location, $timeout) {
        console.log("Meeting Controller initialized");
        $scope.meeting = {
            name: "Placeholder Meeting"
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

