swagger generate server -f haproxy_spec.yaml \
    -A "Data Plane" \
    -t /home/mjuraga/projects/haproxytech/ \
    --existing-models github.com/haproxytech/client-native/v2/models \
    --exclude-main \
    --skip-models \
    -s dataplaneapi \
    --tags=Discovery \
    --tags=ServiceDiscovery \
    --tags=Information \
    --tags=Specification \
    --tags=SpecificationOpenapiv3 \
    --tags=Transactions \
    --tags=Sites \
    --tags=Stats \
    --tags=Global \
    --tags=Frontend \
    --tags=Backend \
    --tags=Bind \
    --tags=Server \
    --tags=Configuration \
    --tags=HTTPRequestRule \
    --tags=HTTPResponseRule \
    --tags=BackendSwitchingRule \
    --tags=ServerSwitchingRule \
    --tags=TCPResponseRule \
    --tags=TCPRequestRule \
    --tags=Filter \
    --tags=StickRule \
    --tags=LogTarget \
    --tags=Reloads \
    --tags=ACL \
    --tags=Defaults \
    --tags=StickTable \
    --tags=Maps \
    --tags=Nameserver \
    --tags=Cluster \
    --tags=Peer \
    --tags=PeerEntry \
    --tags=Resolver \
    --tags=Spoe \
    --tags=SpoeTransactions \
    --tags=Storage \
    --tags="ACL Runtime" \
    --tags=ServerTemplate \
    -r ../copyright.txt
