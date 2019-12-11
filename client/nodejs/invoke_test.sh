node invoke.js dc AddURL D100010001 https://test.iot.org/voice00001.mp3
# node invoke.js dc GetURL D100010001

# node invoke.js pc AddPolicy '{"AS":{"userId":"13800010002","role":"u1","group":"g1"},"AO":{"deviceId":"D100010001","MAC":"00:11:22:33:44:55"},"AP":1,"AE":{"createdTime":1575468182,"endTime":1576468182,"allowedIP":"*.*.*.*"}}'
# node invoke.js pc QueryPolicy da1f08aa8ee250a65057c354b6952e2baf6ae3c9505cb9ca3e5ab4138e56b9d1

# node invoke.js pc AddPolicy '{"AS":{"userId":"13800010001","role":"u1","group":"g1"},"AO":{"deviceId":"D100010001","MAC":"00:11:22:33:44:55"},"AP":1,"AE":{"createdTime":1575468182,"endTime":1576468182,"allowedIP":"*.*.*.*"}}'
# node invoke.js pc QueryPolicy 40db810e4ccb4cc1f3d5bc5803fb61e863cf05ea7fc2f63165599ef53adf5623

# node invoke.js ac CheckAccess '{"AS":{"userId":"13800010002","role":"u1","group":"g1"},"AO":{"deviceId":"D100010001","MAC":"00:11:22:33:44:55"}}'
# node invoke.js ac CheckAccess '{"AS":{"userId":"13800010001","role":"u1","group":"g1"},"AO":{"deviceId":"D100010001","MAC":"00:11:22:33:44:55"}}'