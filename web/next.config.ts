import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    distDir: "build",
    output: "standalone",
    rewrites: async () => {
    	return [
     		{
       			source: "/v1/submissions/1/status/:statustype",
          		destination: "http://localhost:8082/v1/submissions/:statustype"
       		}
     	]
    }
};

export default nextConfig;
