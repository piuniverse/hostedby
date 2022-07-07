conn = new Mongo();
db = conn.getDB("ips");

//Sample IP
db.col_ips.insert(
    {
        _id:            "52.93.153.170",
        ip_address:     "52.93.153.170", 
        url:            "https://ip-ranges.amazonaws.com/ip-ranges.json",
        cloudplatform:  "aws"
    }
 );

 db.col_ips.insert(
    {
        _id:            "34.137.0.2",
        ip_address:     "34.137.0.2", 
        url:            "https://www.gstatic.com/ipranges/cloud.json",
        cloudplatform:  "google"
    }
 );

