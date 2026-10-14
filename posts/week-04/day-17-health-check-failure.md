# Day 17 posts

## LinkedIn

Day 17 of my cloud engineering journey: I practiced diagnosing a failed health probe using my Go service locally.

The correct /healthz path returned HTTP 200 and ok. An incorrect path returned 404 while the application stayed running. Correcting the probe restored HTTP 200 and ok without changing the application.

The lesson: a failed probe does not always mean the process is down. Replacing instances will not repair an incorrect health-check configuration shared by every replacement.

I also documented five failure scenarios, separating detection, automatic recovery, and human repair. The local probe results were tested; the other scenarios were documented analysis, supported by my earlier AWS instance-replacement lab.

I stopped the service, verified port 8080 was released, and created no AWS resources.

#AWS #CloudEngineering #Go #Troubleshooting #LearningInPublic

## X / Threads

Day 17: Go health probe returned 200/ok; a wrong path returned 404 while the app stayed running. Correcting the probe restored success. Replacing instances cannot fix shared bad probe configuration. Local service stopped; no AWS resources. #CloudEngineering
