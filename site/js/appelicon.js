angular.module('apelliconApp', [])
    .controller('SearchController', ['$scope', '$http', function ($scope, $http) {
        $scope.searchText = "";
        $scope.searchResponse = "waiting";

        $scope.updateResults = function () {
            var query = '/api/interfaces?query=' + $scope.searchText;
            console.log("Search text is now " + query);

            $http({
                method: 'GET',
                url: query
            }).then(function successCallback(response) {
                $scope.searchResponse = response.data;
            }, function errorCallback(response) {
                console.log("error from API " + response)
                // called asynchronously if an error occurs
                // or server returns response with an error status.
            });
        };

        console.log("Apellicon App Controller Loaded");
    }]);