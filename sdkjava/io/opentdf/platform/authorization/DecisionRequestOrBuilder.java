// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: authorization/authorization.proto

// Protobuf Java Version: 3.25.3
package io.opentdf.platform.authorization;

public interface DecisionRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:authorization.DecisionRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>repeated .authorization.Action actions = 1 [json_name = "actions"];</code>
   */
  java.util.List<io.opentdf.platform.authorization.Action> 
      getActionsList();
  /**
   * <code>repeated .authorization.Action actions = 1 [json_name = "actions"];</code>
   */
  io.opentdf.platform.authorization.Action getActions(int index);
  /**
   * <code>repeated .authorization.Action actions = 1 [json_name = "actions"];</code>
   */
  int getActionsCount();
  /**
   * <code>repeated .authorization.Action actions = 1 [json_name = "actions"];</code>
   */
  java.util.List<? extends io.opentdf.platform.authorization.ActionOrBuilder> 
      getActionsOrBuilderList();
  /**
   * <code>repeated .authorization.Action actions = 1 [json_name = "actions"];</code>
   */
  io.opentdf.platform.authorization.ActionOrBuilder getActionsOrBuilder(
      int index);

  /**
   * <code>repeated .authorization.EntityChain entity_chains = 2 [json_name = "entityChains"];</code>
   */
  java.util.List<io.opentdf.platform.authorization.EntityChain> 
      getEntityChainsList();
  /**
   * <code>repeated .authorization.EntityChain entity_chains = 2 [json_name = "entityChains"];</code>
   */
  io.opentdf.platform.authorization.EntityChain getEntityChains(int index);
  /**
   * <code>repeated .authorization.EntityChain entity_chains = 2 [json_name = "entityChains"];</code>
   */
  int getEntityChainsCount();
  /**
   * <code>repeated .authorization.EntityChain entity_chains = 2 [json_name = "entityChains"];</code>
   */
  java.util.List<? extends io.opentdf.platform.authorization.EntityChainOrBuilder> 
      getEntityChainsOrBuilderList();
  /**
   * <code>repeated .authorization.EntityChain entity_chains = 2 [json_name = "entityChains"];</code>
   */
  io.opentdf.platform.authorization.EntityChainOrBuilder getEntityChainsOrBuilder(
      int index);

  /**
   * <code>repeated .authorization.ResourceAttributes resource_attributes = 3 [json_name = "resourceAttributes"];</code>
   */
  java.util.List<io.opentdf.platform.authorization.ResourceAttributes> 
      getResourceAttributesList();
  /**
   * <code>repeated .authorization.ResourceAttributes resource_attributes = 3 [json_name = "resourceAttributes"];</code>
   */
  io.opentdf.platform.authorization.ResourceAttributes getResourceAttributes(int index);
  /**
   * <code>repeated .authorization.ResourceAttributes resource_attributes = 3 [json_name = "resourceAttributes"];</code>
   */
  int getResourceAttributesCount();
  /**
   * <code>repeated .authorization.ResourceAttributes resource_attributes = 3 [json_name = "resourceAttributes"];</code>
   */
  java.util.List<? extends io.opentdf.platform.authorization.ResourceAttributesOrBuilder> 
      getResourceAttributesOrBuilderList();
  /**
   * <code>repeated .authorization.ResourceAttributes resource_attributes = 3 [json_name = "resourceAttributes"];</code>
   */
  io.opentdf.platform.authorization.ResourceAttributesOrBuilder getResourceAttributesOrBuilder(
      int index);
}