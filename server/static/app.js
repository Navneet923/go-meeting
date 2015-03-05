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
            }).when("/actions", {
                templateUrl: "partials/action-list.html",
                controller:  "ActionListController"
            }).when("/new_meeting", {
                templateUrl: "partials/new-meeting.html",
                controller:  "NewMeetingController"
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

        $scope.activeTab = 'meetings';
        $scope.changeTab = function(targetTabName) {
            $scope.activeTab = targetTabName;
            $location.url("" + $scope.activeTab);
        }

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

/******************************************************************************\
Function:
    ActionListController 

Dependencies:
    ...

Description:
    Application controller responsible for querying and displaying action items 
\******************************************************************************/
app.controller("ActionListController", ["$scope", "$location", "$timeout",

    function($scope, $location, $timeout) {
        console.log("Action List Controller initialized");
    }]

);

/******************************************************************************\
Function:
    NewMeetingController 

Dependencies:
    ...

Description:
    Application controller responsible for updating / creating a new meeting 
\******************************************************************************/
app.controller("NewMeetingController", ["$scope", "$location", "$timeout",

    function($scope, $location, $timeout) {
        console.log("New Meeting Controller initialized");
        $scope.meeting = {};

        $scope.submitNewMeeting = function() {
            console.log("TODO: Submit meeting request to server");
            console.log($scope.meeting);
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

