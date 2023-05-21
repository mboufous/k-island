// Components
export type NamespaceDetailsParams = {
  name: string;
};

// Common
export interface Spec {
  blocked: boolean;
  branch: string;
  commit: string;
  image_tag: string;
  project: string;
  sha256: string;
}

// Namespaces List

export type NamespaceItem = {
  name: string;
  spec: Spec;
  blocked: boolean;
  status: string;
};

export type NamespacesHTTPResponse = NamespaceItem[];

// Namespace Details

export interface NamespaceDetailsHTTPResponse {
  id: string;
  env_info: EnvInfo;
  ingresses: Record<string, Ingress>;
  pods: Record<string, Pod>;
}

export interface EnvInfo {
  status: EnvInfoStatus;
  spec: Spec;
}

export interface EnvInfoStatus {
  general: boolean;
  lastCheckTimestamp: Date;
  lastUpdateTimestamp: Date;
  phase: string;
  resources: Resource[];
  warnings: Warnings;
}

export interface Resource {
  name: string;
  replica_status: ReplicaStatus;
  status: boolean;
}

export interface ReplicaStatus {
  available_replicas?: number;
  ready_replicas: number;
  replicas: number;
  updated_replicas: number;
  current_replicas?: number;
}

export interface Warnings {
  pod: Pod;
}

export interface Pod {
  namespace: string;
  phase: string;
  containers: Record<string, Container>;
  kibana_url: string;
  message: Message[];
}

type Container = {
  image: string;
};

export interface Message {
  firstTimestamp: Date;
  lastTimestamp: Date;
  message: string;
  pod_name: string;
  reason: string;
}

export interface Ingress {
  rules: Rule[];
}

export interface Rule {
  host: string;
  http: HTTP;
}

export interface HTTP {
  paths: Path[];
}

export interface Path {
  backend: Backend;
  path: string;
}

export interface Backend {
  resource: null;
  service: Service;
}

export interface Service {
  name: string;
  port: Port;
}

export interface Port {
  name: null | string;
  number: number | null;
}
