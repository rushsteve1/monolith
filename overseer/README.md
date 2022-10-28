# Overseer

The Overseer brings together all the different modules and starts them.
It then acts as a supervisor using `suture` and generally provides a backbone
for the services. It also provides the RPC interface that [`socon`](../socon/)
communicates with
