## tableflip example
### What is [tableflip](https://github.com/cloudflare/tableflip)
> It is sometimes useful to update the running code and / or configuration of a network service, without disrupting existing connections. Usually, this is achieved by starting a new process, somehow transferring clients to it and then exiting the old process.

> There are many ways to implement graceful upgrades. They vary wildly in the trade-offs they make, and how much control they afford the user. This library has the following goals:
>  - No old code keeps running after a successful upgrade
>  - The new process has a grace period for performing initialisation
>  - Crashing during initialisation is OK
>  - Only a single upgrade is ever run in parallel
>  - tableflip works on Linux and macOS.

### Steps
1. Build v0.0.1 demo.
   ```bash
   go build -o demo main.go
   ```
2. Run the demo and create a get request to `127.0.0.1:8080/version`, here is the output:
   ```bash
   [PID: 123] v0.0.1
   ```
3. Prepare a new version. change the main.go, let the version be "v0.0.2" and rebuild it.
   ```bash
   go build -o demo main.go
   ```
4. Now, kill the old one !
   ```bash
   kill -s HUP 123
   ```
5. Create the request to version api again, but output is changed:
   ```bash
   [PID: 123] v0.0.2
   ```

The client is completely immune to server upgrades and reboots, and our application updates gracefully!
