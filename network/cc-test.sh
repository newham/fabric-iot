# upgrade
# shell|action|cc_name|cc_version|cc_src|fname|args
./cc.sh upgrade pc 1.1 go/dc Synchro

# install
# shell|action|cc_name|cc_version|cc_src|fname|args
./cc.sh install pc 1.0 go/ac Synchro

# invoke
# shell|action|cc_name|cc_version|cc_src|fname|args
# add policy
./cc.sh invoke pc 1.0 go/pc AddPolicy '"{\"AS\":{\"userId\":\"13800010001\",\"role\":\"u1\",\"group\":\"g1\"},\"AO\":{\"deviceId\":\"D100010001\",\"MAC\":\"00:11:22:33:44:55\"}}"'
# query policy
./cc.sh invoke pc 1.0 go/pc QueryPolicy '"40db810e4ccb4cc1f3d5bc5803fb61e863cf05ea7fc2f63165599ef53adf5623"'