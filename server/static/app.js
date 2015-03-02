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

