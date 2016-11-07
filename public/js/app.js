// Defining angularjs application.
    var postApp = angular.module('postApp', []);
    // Controller function and passing $http service and $scope var.
    postApp.controller('postController', function($scope, $http) {
      // create a blank object to handle form data.
        $scope.name = {};
      // calling our submit function.
        $scope.submitForm = function() {
        // Posting data to GO file
        $http({
          method  : 'POST',
          url     : 'webapp.go',
          data    : $scope.name, //forms user object
          headers : {'Content-Type': 'application/x-www-form-urlencoded'} 
         })
          .success(function(data) {
            $scope.name = data.message;
          });
        };
    });