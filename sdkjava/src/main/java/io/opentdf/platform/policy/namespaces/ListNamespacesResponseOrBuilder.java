// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: policy/namespaces/namespaces.proto

// Protobuf Java Version: 3.25.3
package io.opentdf.platform.policy.namespaces;

public interface ListNamespacesResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:policy.namespaces.ListNamespacesResponse)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>repeated .policy.namespaces.Namespace namespaces = 1 [json_name = "namespaces"];</code>
   */
  java.util.List<io.opentdf.platform.policy.namespaces.Namespace> 
      getNamespacesList();
  /**
   * <code>repeated .policy.namespaces.Namespace namespaces = 1 [json_name = "namespaces"];</code>
   */
  io.opentdf.platform.policy.namespaces.Namespace getNamespaces(int index);
  /**
   * <code>repeated .policy.namespaces.Namespace namespaces = 1 [json_name = "namespaces"];</code>
   */
  int getNamespacesCount();
  /**
   * <code>repeated .policy.namespaces.Namespace namespaces = 1 [json_name = "namespaces"];</code>
   */
  java.util.List<? extends io.opentdf.platform.policy.namespaces.NamespaceOrBuilder> 
      getNamespacesOrBuilderList();
  /**
   * <code>repeated .policy.namespaces.Namespace namespaces = 1 [json_name = "namespaces"];</code>
   */
  io.opentdf.platform.policy.namespaces.NamespaceOrBuilder getNamespacesOrBuilder(
      int index);
}