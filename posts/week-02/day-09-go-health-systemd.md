# Day 9 Build-in-Public Draft — Go and systemd on EC2

## LinkedIn

Day 9 of my AWS cloud engineering journey connected application development with Linux service operations.

I built a dependency-free Go endpoint that returns HTTP 200 with `ok` at `127.0.0.1:8080/healthz`. I compiled it on a short-lived Amazon Linux 2023 EC2 instance and installed it as a non-root `systemd` service with `Restart=on-failure`.

I checked four different layers:

- `systemctl` confirmed the service process was active.
- `ss` confirmed the loopback listener existed on port 8080.
- `curl` confirmed the application returned the expected HTTP response.
- `journalctl` showed the service and application startup events.

I then sent SIGKILL to the main process. Within seconds, `systemd` restarted the service and the health endpoint returned `ok` again.

The instance had no inbound security-group rules and no SSH key. After the test, I terminated it and confirmed that no Day 9 EBS volume remained.

My takeaway: “the server is up” is not one check. Process state, socket state, HTTP behavior, and logs answer different questions, and reliable troubleshooting uses all four.

#AWS #CloudEngineering #Go #Linux #EC2 #systemd

## X

AWS Day 9: ran a Go health endpoint as a non-root systemd service on Amazon Linux 2023. Verified process, socket, HTTP response, and logs, then killed the process and watched systemd restore it. Instance terminated; no EBS volume remained. #AWS #Go #Linux
