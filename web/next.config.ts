import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    distDir: "build",
    output: "standalone",
    rewrites: async () => {
    	return [
     		{
       			source: "/v1/submissions/:submissionid/status/:statustype",
          		destination: "http://mapsservice:8082/v1/submissions/:submissionid/status/:statustype"
       		}
     	]
    }
};

export default nextConfig;
