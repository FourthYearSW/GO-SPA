var appname = angular.module('appname', []);

appname.controller('postCtrl', function($scope){

    $scope.name = {};

    $scope.firstName = { text: 'christy' };
    $scope.surname = 'madden';

});

// URL is localhost and port number, then the path of the file data is being sent to