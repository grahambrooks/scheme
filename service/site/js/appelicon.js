angular.module('apelliconApp', [])
    .controller('SearchController', ['$scope', '$http', function ($scope, $http) {
        $scope.searchText = "";
        $scope.searchResponse = "waiting";
        $scope.resultCount = 0

        $scope.updateResults = function () {
            var query = '/api/search?query=' + $scope.searchText;

            $http({
                method: 'GET',
                url: query
            }).then(function successCallback(response) {
                $scope.searchResponse = response.data;
                $scope.resultCount = $scope.searchResponse.hits.total.value
            }, function errorCallback(response) {
                console.log("error from API " + response)
                // called asynchronously if an error occurs
                // or server returns response with an error status.
            });
        };

        $scope.updateResults()
    }]);