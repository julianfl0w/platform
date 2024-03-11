// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: policy/namespaces/namespaces.proto

// Protobuf Java Version: 3.25.3
package io.opentdf.platform.policy.namespaces;

public interface NamespaceOrBuilder extends
    // @@protoc_insertion_point(interface_extends:policy.namespaces.Namespace)
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
   * <code>string name = 2 [json_name = "name"];</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <pre>
   * used to partition Attribute Definitions, support by namespace AuthN and enable federation
   * </pre>
   *
   * <code>string name = 2 [json_name = "name"];</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>string fqn = 3 [json_name = "fqn"];</code>
   * @return The fqn.
   */
  java.lang.String getFqn();
  /**
   * <code>string fqn = 3 [json_name = "fqn"];</code>
   * @return The bytes for fqn.
   */
  com.google.protobuf.ByteString
      getFqnBytes();

  /**
   * <pre>
   * active by default until explicitly deactivated
   * </pre>
   *
   * <code>.google.protobuf.BoolValue active = 4 [json_name = "active"];</code>
   * @return Whether the active field is set.
   */
  boolean hasActive();
  /**
   * <pre>
   * active by default until explicitly deactivated
   * </pre>
   *
   * <code>.google.protobuf.BoolValue active = 4 [json_name = "active"];</code>
   * @return The active.
   */
  com.google.protobuf.BoolValue getActive();
  /**
   * <pre>
   * active by default until explicitly deactivated
   * </pre>
   *
   * <code>.google.protobuf.BoolValue active = 4 [json_name = "active"];</code>
   */
  com.google.protobuf.BoolValueOrBuilder getActiveOrBuilder();

  /**
   * <code>.common.Metadata metadata = 5 [json_name = "metadata"];</code>
   * @return Whether the metadata field is set.
   */
  boolean hasMetadata();
  /**
   * <code>.common.Metadata metadata = 5 [json_name = "metadata"];</code>
   * @return The metadata.
   */
  io.opentdf.platform.common.Metadata getMetadata();
  /**
   * <code>.common.Metadata metadata = 5 [json_name = "metadata"];</code>
   */
  io.opentdf.platform.common.MetadataOrBuilder getMetadataOrBuilder();
}