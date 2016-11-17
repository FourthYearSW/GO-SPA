var appname = angular.module('appname', []);

appname.controller('appCtrl', ['$scope',
  function($scope) {
    $scope.name = { text: 'christy' };
    $scope.surname = 'madden';
}]);