// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: namespaces/namespaces.proto

// Protobuf Java Version: 3.25.2
package com.namespaces;

public interface NamespaceOrBuilder extends
    // @@protoc_insertion_point(interface_extends:namespaces.Namespace)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * generated uuid in database
   * </pre>
   *
   * <code>string id = 1 [json_name = "id"];</code>
   * @return The id.
   */
  java.lang.String getId();
  /**
   * <pre>
   * generated uuid in database
   * </pre>
   *
   * <code>string id = 1 [json_name = "id"];</code>
   * @return The bytes for id.
   */
  com.google.protobuf.ByteString
      getIdBytes();

  /**
   * <pre>
   * used to partition Attribute Definitions, support by namespace AuthN and enable federation
   * </pre>
   *
   * <code>string name = 5 [json_name = "name", (.buf.validate.field) = { ... }</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <pre>
   * used to partition Attribute Definitions, support by namespace AuthN and enable federation
   * </pre>
   *
   * <code>string name = 5 [json_name = "name", (.buf.validate.field) = { ... }</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();
}