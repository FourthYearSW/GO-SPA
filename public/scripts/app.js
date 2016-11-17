var appname = angular.module('appname', []);

appname.controller('appCtrl', ['$scope',
  function($scope) {
    $scope.name = { text: 'christy' };
    $scope.surname = 'madden';
}]);

// URL is localhost and port number, then the path of the file data is being sent to