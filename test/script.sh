curl -vX POST  http://localhost:8080/api/v1alpha1/add_schedule -d @schedule.json \
--header "Content-Type: application/json"
 curl  http://localhost:8080/api/v1alpha1/get_schedule\?schedule=test --header "Content-Type: application/json"

