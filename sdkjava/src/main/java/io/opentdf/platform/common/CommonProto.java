// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: common/common.proto

// Protobuf Java Version: 3.25.3
package io.opentdf.platform.common;

public final class CommonProto {
  private CommonProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_common_Metadata_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_common_Metadata_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_common_Metadata_LabelsEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_common_Metadata_LabelsEntry_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_common_MetadataMutable_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_common_MetadataMutable_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_common_MetadataMutable_LabelsEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_common_MetadataMutable_LabelsEntry_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\023common/common.proto\022\006common\032\037google/pr" +
      "otobuf/timestamp.proto\"\361\001\n\010Metadata\0229\n\nc" +
      "reated_at\030\001 \001(\0132\032.google.protobuf.Timest" +
      "ampR\tcreatedAt\0229\n\nupdated_at\030\002 \001(\0132\032.goo" +
      "gle.protobuf.TimestampR\tupdatedAt\0224\n\006lab" +
      "els\030\003 \003(\0132\034.common.Metadata.LabelsEntryR" +
      "\006labels\0329\n\013LabelsEntry\022\020\n\003key\030\001 \001(\tR\003key" +
      "\022\024\n\005value\030\002 \001(\tR\005value:\0028\001\"\211\001\n\017MetadataM" +
      "utable\022;\n\006labels\030\003 \003(\0132#.common.Metadata" +
      "Mutable.LabelsEntryR\006labels\0329\n\013LabelsEnt" +
      "ry\022\020\n\003key\030\001 \001(\tR\003key\022\024\n\005value\030\002 \001(\tR\005val" +
      "ue:\0028\001*}\n\022MetadataUpdateEnum\022$\n METADATA" +
      "_UPDATE_ENUM_UNSPECIFIED\020\000\022\037\n\033METADATA_U" +
      "PDATE_ENUM_EXTEND\020\001\022 \n\034METADATA_UPDATE_E" +
      "NUM_REPLACE\020\002*\215\001\n\017ActiveStateEnum\022!\n\035ACT" +
      "IVE_STATE_ENUM_UNSPECIFIED\020\000\022\034\n\030ACTIVE_S" +
      "TATE_ENUM_ACTIVE\020\001\022\036\n\032ACTIVE_STATE_ENUM_" +
      "INACTIVE\020\002\022\031\n\025ACTIVE_STATE_ENUM_ANY\020\003B\221\001" +
      "\n\032io.opentdf.platform.commonB\013CommonProt" +
      "oP\001Z.github.com/opentdf/platform/protoco" +
      "l/go/common\242\002\003CXX\252\002\006Common\312\002\006Common\342\002\022Co" +
      "mmon\\GPBMetadata\352\002\006Commonb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.TimestampProto.getDescriptor(),
        });
    internal_static_common_Metadata_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_common_Metadata_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_common_Metadata_descriptor,
        new java.lang.String[] { "CreatedAt", "UpdatedAt", "Labels", });
    internal_static_common_Metadata_LabelsEntry_descriptor =
      internal_static_common_Metadata_descriptor.getNestedTypes().get(0);
    internal_static_common_Metadata_LabelsEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_common_Metadata_LabelsEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    internal_static_common_MetadataMutable_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_common_MetadataMutable_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_common_MetadataMutable_descriptor,
        new java.lang.String[] { "Labels", });
    internal_static_common_MetadataMutable_LabelsEntry_descriptor =
      internal_static_common_MetadataMutable_descriptor.getNestedTypes().get(0);
    internal_static_common_MetadataMutable_LabelsEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_common_MetadataMutable_LabelsEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    com.google.protobuf.TimestampProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}