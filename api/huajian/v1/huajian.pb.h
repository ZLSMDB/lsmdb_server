// Generated by the protocol buffer compiler.  DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: huajian/v1/huajian.proto
// Protobuf C++ Version: 5.27.0

#ifndef GOOGLE_PROTOBUF_INCLUDED_huajian_2fv1_2fhuajian_2eproto_2epb_2eh
#define GOOGLE_PROTOBUF_INCLUDED_huajian_2fv1_2fhuajian_2eproto_2epb_2eh

#include <limits>
#include <string>
#include <type_traits>
#include <utility>

#include "google/protobuf/runtime_version.h"
#if PROTOBUF_VERSION != 5027000
#error "Protobuf C++ gencode is built with an incompatible version of"
#error "Protobuf C++ headers/runtime. See"
#error "https://protobuf.dev/support/cross-version-runtime-guarantee/#cpp"
#endif
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/arena.h"
#include "google/protobuf/arenastring.h"
#include "google/protobuf/generated_message_tctable_decl.h"
#include "google/protobuf/generated_message_util.h"
#include "google/protobuf/metadata_lite.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/message.h"
#include "google/protobuf/repeated_field.h"  // IWYU pragma: export
#include "google/protobuf/extension_set.h"  // IWYU pragma: export
#include "google/protobuf/unknown_field_set.h"
#include "google/api/annotations.pb.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"

#define PROTOBUF_INTERNAL_EXPORT_huajian_2fv1_2fhuajian_2eproto

namespace google {
namespace protobuf {
namespace internal {
class AnyMetadata;
}  // namespace internal
}  // namespace protobuf
}  // namespace google

// Internal implementation detail -- do not use these members.
struct TableStruct_huajian_2fv1_2fhuajian_2eproto {
  static const ::uint32_t offsets[];
};
extern const ::google::protobuf::internal::DescriptorTable
    descriptor_table_huajian_2fv1_2fhuajian_2eproto;
namespace helloworld {
namespace v1 {
class GetModelReply;
struct GetModelReplyDefaultTypeInternal;
extern GetModelReplyDefaultTypeInternal _GetModelReply_default_instance_;
class GetModelRequest;
struct GetModelRequestDefaultTypeInternal;
extern GetModelRequestDefaultTypeInternal _GetModelRequest_default_instance_;
}  // namespace v1
}  // namespace helloworld
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google

namespace helloworld {
namespace v1 {

// ===================================================================


// -------------------------------------------------------------------

class GetModelRequest final : public ::google::protobuf::Message
/* @@protoc_insertion_point(class_definition:helloworld.v1.GetModelRequest) */ {
 public:
  inline GetModelRequest() : GetModelRequest(nullptr) {}
  ~GetModelRequest() override;
  template <typename = void>
  explicit PROTOBUF_CONSTEXPR GetModelRequest(
      ::google::protobuf::internal::ConstantInitialized);

  inline GetModelRequest(const GetModelRequest& from) : GetModelRequest(nullptr, from) {}
  inline GetModelRequest(GetModelRequest&& from) noexcept
      : GetModelRequest(nullptr, std::move(from)) {}
  inline GetModelRequest& operator=(const GetModelRequest& from) {
    CopyFrom(from);
    return *this;
  }
  inline GetModelRequest& operator=(GetModelRequest&& from) noexcept {
    if (this == &from) return *this;
    if (GetArena() == from.GetArena()
#ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetArena() != nullptr
#endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const
      ABSL_ATTRIBUTE_LIFETIME_BOUND {
    return _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance);
  }
  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields()
      ABSL_ATTRIBUTE_LIFETIME_BOUND {
    return _internal_metadata_.mutable_unknown_fields<::google::protobuf::UnknownFieldSet>();
  }

  static const ::google::protobuf::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::google::protobuf::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::google::protobuf::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const GetModelRequest& default_instance() {
    return *internal_default_instance();
  }
  static inline const GetModelRequest* internal_default_instance() {
    return reinterpret_cast<const GetModelRequest*>(
        &_GetModelRequest_default_instance_);
  }
  static constexpr int kIndexInFileMessages = 0;
  friend void swap(GetModelRequest& a, GetModelRequest& b) { a.Swap(&b); }
  inline void Swap(GetModelRequest* other) {
    if (other == this) return;
#ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetArena() != nullptr && GetArena() == other->GetArena()) {
#else   // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetArena() == other->GetArena()) {
#endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::google::protobuf::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(GetModelRequest* other) {
    if (other == this) return;
    ABSL_DCHECK(GetArena() == other->GetArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  GetModelRequest* New(::google::protobuf::Arena* arena = nullptr) const final {
    return ::google::protobuf::Message::DefaultConstruct<GetModelRequest>(arena);
  }
  using ::google::protobuf::Message::CopyFrom;
  void CopyFrom(const GetModelRequest& from);
  using ::google::protobuf::Message::MergeFrom;
  void MergeFrom(const GetModelRequest& from) { GetModelRequest::MergeImpl(*this, from); }

  private:
  static void MergeImpl(
      ::google::protobuf::MessageLite& to_msg,
      const ::google::protobuf::MessageLite& from_msg);

  public:
  bool IsInitialized() const {
    return true;
  }
  ABSL_ATTRIBUTE_REINITIALIZES void Clear() final;
  ::size_t ByteSizeLong() const final;
  ::uint8_t* _InternalSerialize(
      ::uint8_t* target,
      ::google::protobuf::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::google::protobuf::Arena* arena);
  void SharedDtor();
  void InternalSwap(GetModelRequest* other);
 private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() { return "helloworld.v1.GetModelRequest"; }

 protected:
  explicit GetModelRequest(::google::protobuf::Arena* arena);
  GetModelRequest(::google::protobuf::Arena* arena, const GetModelRequest& from);
  GetModelRequest(::google::protobuf::Arena* arena, GetModelRequest&& from) noexcept
      : GetModelRequest(arena) {
    *this = ::std::move(from);
  }
  const ::google::protobuf::Message::ClassData* GetClassData() const final;

 public:
  ::google::protobuf::Metadata GetMetadata() const;
  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------
  enum : int {
    kDbNameFieldNumber = 1,
    kKeyFieldNumber = 2,
  };
  // string db_name = 1;
  void clear_db_name() ;
  const std::string& db_name() const;
  template <typename Arg_ = const std::string&, typename... Args_>
  void set_db_name(Arg_&& arg, Args_... args);
  std::string* mutable_db_name();
  PROTOBUF_NODISCARD std::string* release_db_name();
  void set_allocated_db_name(std::string* value);

  private:
  const std::string& _internal_db_name() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_db_name(
      const std::string& value);
  std::string* _internal_mutable_db_name();

  public:
  // string key = 2;
  void clear_key() ;
  const std::string& key() const;
  template <typename Arg_ = const std::string&, typename... Args_>
  void set_key(Arg_&& arg, Args_... args);
  std::string* mutable_key();
  PROTOBUF_NODISCARD std::string* release_key();
  void set_allocated_key(std::string* value);

  private:
  const std::string& _internal_key() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_key(
      const std::string& value);
  std::string* _internal_mutable_key();

  public:
  // @@protoc_insertion_point(class_scope:helloworld.v1.GetModelRequest)
 private:
  class _Internal;
  friend class ::google::protobuf::internal::TcParser;
  static const ::google::protobuf::internal::TcParseTable<
      1, 2, 0,
      48, 2>
      _table_;

  static constexpr const void* _raw_default_instance_ =
      &_GetModelRequest_default_instance_;

  friend class ::google::protobuf::MessageLite;
  friend class ::google::protobuf::Arena;
  template <typename T>
  friend class ::google::protobuf::Arena::InternalHelper;
  using InternalArenaConstructable_ = void;
  using DestructorSkippable_ = void;
  struct Impl_ {
    inline explicit constexpr Impl_(
        ::google::protobuf::internal::ConstantInitialized) noexcept;
    inline explicit Impl_(::google::protobuf::internal::InternalVisibility visibility,
                          ::google::protobuf::Arena* arena);
    inline explicit Impl_(::google::protobuf::internal::InternalVisibility visibility,
                          ::google::protobuf::Arena* arena, const Impl_& from,
                          const GetModelRequest& from_msg);
    ::google::protobuf::internal::ArenaStringPtr db_name_;
    ::google::protobuf::internal::ArenaStringPtr key_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    PROTOBUF_TSAN_DECLARE_MEMBER
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_huajian_2fv1_2fhuajian_2eproto;
};
// -------------------------------------------------------------------

class GetModelReply final : public ::google::protobuf::Message
/* @@protoc_insertion_point(class_definition:helloworld.v1.GetModelReply) */ {
 public:
  inline GetModelReply() : GetModelReply(nullptr) {}
  ~GetModelReply() override;
  template <typename = void>
  explicit PROTOBUF_CONSTEXPR GetModelReply(
      ::google::protobuf::internal::ConstantInitialized);

  inline GetModelReply(const GetModelReply& from) : GetModelReply(nullptr, from) {}
  inline GetModelReply(GetModelReply&& from) noexcept
      : GetModelReply(nullptr, std::move(from)) {}
  inline GetModelReply& operator=(const GetModelReply& from) {
    CopyFrom(from);
    return *this;
  }
  inline GetModelReply& operator=(GetModelReply&& from) noexcept {
    if (this == &from) return *this;
    if (GetArena() == from.GetArena()
#ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetArena() != nullptr
#endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const
      ABSL_ATTRIBUTE_LIFETIME_BOUND {
    return _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance);
  }
  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields()
      ABSL_ATTRIBUTE_LIFETIME_BOUND {
    return _internal_metadata_.mutable_unknown_fields<::google::protobuf::UnknownFieldSet>();
  }

  static const ::google::protobuf::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::google::protobuf::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::google::protobuf::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const GetModelReply& default_instance() {
    return *internal_default_instance();
  }
  static inline const GetModelReply* internal_default_instance() {
    return reinterpret_cast<const GetModelReply*>(
        &_GetModelReply_default_instance_);
  }
  static constexpr int kIndexInFileMessages = 1;
  friend void swap(GetModelReply& a, GetModelReply& b) { a.Swap(&b); }
  inline void Swap(GetModelReply* other) {
    if (other == this) return;
#ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetArena() != nullptr && GetArena() == other->GetArena()) {
#else   // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetArena() == other->GetArena()) {
#endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::google::protobuf::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(GetModelReply* other) {
    if (other == this) return;
    ABSL_DCHECK(GetArena() == other->GetArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  GetModelReply* New(::google::protobuf::Arena* arena = nullptr) const final {
    return ::google::protobuf::Message::DefaultConstruct<GetModelReply>(arena);
  }
  using ::google::protobuf::Message::CopyFrom;
  void CopyFrom(const GetModelReply& from);
  using ::google::protobuf::Message::MergeFrom;
  void MergeFrom(const GetModelReply& from) { GetModelReply::MergeImpl(*this, from); }

  private:
  static void MergeImpl(
      ::google::protobuf::MessageLite& to_msg,
      const ::google::protobuf::MessageLite& from_msg);

  public:
  bool IsInitialized() const {
    return true;
  }
  ABSL_ATTRIBUTE_REINITIALIZES void Clear() final;
  ::size_t ByteSizeLong() const final;
  ::uint8_t* _InternalSerialize(
      ::uint8_t* target,
      ::google::protobuf::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::google::protobuf::Arena* arena);
  void SharedDtor();
  void InternalSwap(GetModelReply* other);
 private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() { return "helloworld.v1.GetModelReply"; }

 protected:
  explicit GetModelReply(::google::protobuf::Arena* arena);
  GetModelReply(::google::protobuf::Arena* arena, const GetModelReply& from);
  GetModelReply(::google::protobuf::Arena* arena, GetModelReply&& from) noexcept
      : GetModelReply(arena) {
    *this = ::std::move(from);
  }
  const ::google::protobuf::Message::ClassData* GetClassData() const final;

 public:
  ::google::protobuf::Metadata GetMetadata() const;
  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------
  enum : int {
    kValueFieldNumber = 1,
  };
  // bytes value = 1;
  void clear_value() ;
  const std::string& value() const;
  template <typename Arg_ = const std::string&, typename... Args_>
  void set_value(Arg_&& arg, Args_... args);
  std::string* mutable_value();
  PROTOBUF_NODISCARD std::string* release_value();
  void set_allocated_value(std::string* value);

  private:
  const std::string& _internal_value() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_value(
      const std::string& value);
  std::string* _internal_mutable_value();

  public:
  // @@protoc_insertion_point(class_scope:helloworld.v1.GetModelReply)
 private:
  class _Internal;
  friend class ::google::protobuf::internal::TcParser;
  static const ::google::protobuf::internal::TcParseTable<
      0, 1, 0,
      0, 2>
      _table_;

  static constexpr const void* _raw_default_instance_ =
      &_GetModelReply_default_instance_;

  friend class ::google::protobuf::MessageLite;
  friend class ::google::protobuf::Arena;
  template <typename T>
  friend class ::google::protobuf::Arena::InternalHelper;
  using InternalArenaConstructable_ = void;
  using DestructorSkippable_ = void;
  struct Impl_ {
    inline explicit constexpr Impl_(
        ::google::protobuf::internal::ConstantInitialized) noexcept;
    inline explicit Impl_(::google::protobuf::internal::InternalVisibility visibility,
                          ::google::protobuf::Arena* arena);
    inline explicit Impl_(::google::protobuf::internal::InternalVisibility visibility,
                          ::google::protobuf::Arena* arena, const Impl_& from,
                          const GetModelReply& from_msg);
    ::google::protobuf::internal::ArenaStringPtr value_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    PROTOBUF_TSAN_DECLARE_MEMBER
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_huajian_2fv1_2fhuajian_2eproto;
};

// ===================================================================




// ===================================================================


#ifdef __GNUC__
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// -------------------------------------------------------------------

// GetModelRequest

// string db_name = 1;
inline void GetModelRequest::clear_db_name() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.db_name_.ClearToEmpty();
}
inline const std::string& GetModelRequest::db_name() const
    ABSL_ATTRIBUTE_LIFETIME_BOUND {
  // @@protoc_insertion_point(field_get:helloworld.v1.GetModelRequest.db_name)
  return _internal_db_name();
}
template <typename Arg_, typename... Args_>
inline PROTOBUF_ALWAYS_INLINE void GetModelRequest::set_db_name(Arg_&& arg,
                                                     Args_... args) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.db_name_.Set(static_cast<Arg_&&>(arg), args..., GetArena());
  // @@protoc_insertion_point(field_set:helloworld.v1.GetModelRequest.db_name)
}
inline std::string* GetModelRequest::mutable_db_name() ABSL_ATTRIBUTE_LIFETIME_BOUND {
  std::string* _s = _internal_mutable_db_name();
  // @@protoc_insertion_point(field_mutable:helloworld.v1.GetModelRequest.db_name)
  return _s;
}
inline const std::string& GetModelRequest::_internal_db_name() const {
  ::google::protobuf::internal::TSanRead(&_impl_);
  return _impl_.db_name_.Get();
}
inline void GetModelRequest::_internal_set_db_name(const std::string& value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.db_name_.Set(value, GetArena());
}
inline std::string* GetModelRequest::_internal_mutable_db_name() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  return _impl_.db_name_.Mutable( GetArena());
}
inline std::string* GetModelRequest::release_db_name() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  // @@protoc_insertion_point(field_release:helloworld.v1.GetModelRequest.db_name)
  return _impl_.db_name_.Release();
}
inline void GetModelRequest::set_allocated_db_name(std::string* value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.db_name_.SetAllocated(value, GetArena());
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        if (_impl_.db_name_.IsDefault()) {
          _impl_.db_name_.Set("", GetArena());
        }
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:helloworld.v1.GetModelRequest.db_name)
}

// string key = 2;
inline void GetModelRequest::clear_key() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.key_.ClearToEmpty();
}
inline const std::string& GetModelRequest::key() const
    ABSL_ATTRIBUTE_LIFETIME_BOUND {
  // @@protoc_insertion_point(field_get:helloworld.v1.GetModelRequest.key)
  return _internal_key();
}
template <typename Arg_, typename... Args_>
inline PROTOBUF_ALWAYS_INLINE void GetModelRequest::set_key(Arg_&& arg,
                                                     Args_... args) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.key_.Set(static_cast<Arg_&&>(arg), args..., GetArena());
  // @@protoc_insertion_point(field_set:helloworld.v1.GetModelRequest.key)
}
inline std::string* GetModelRequest::mutable_key() ABSL_ATTRIBUTE_LIFETIME_BOUND {
  std::string* _s = _internal_mutable_key();
  // @@protoc_insertion_point(field_mutable:helloworld.v1.GetModelRequest.key)
  return _s;
}
inline const std::string& GetModelRequest::_internal_key() const {
  ::google::protobuf::internal::TSanRead(&_impl_);
  return _impl_.key_.Get();
}
inline void GetModelRequest::_internal_set_key(const std::string& value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.key_.Set(value, GetArena());
}
inline std::string* GetModelRequest::_internal_mutable_key() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  return _impl_.key_.Mutable( GetArena());
}
inline std::string* GetModelRequest::release_key() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  // @@protoc_insertion_point(field_release:helloworld.v1.GetModelRequest.key)
  return _impl_.key_.Release();
}
inline void GetModelRequest::set_allocated_key(std::string* value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.key_.SetAllocated(value, GetArena());
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        if (_impl_.key_.IsDefault()) {
          _impl_.key_.Set("", GetArena());
        }
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:helloworld.v1.GetModelRequest.key)
}

// -------------------------------------------------------------------

// GetModelReply

// bytes value = 1;
inline void GetModelReply::clear_value() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.value_.ClearToEmpty();
}
inline const std::string& GetModelReply::value() const
    ABSL_ATTRIBUTE_LIFETIME_BOUND {
  // @@protoc_insertion_point(field_get:helloworld.v1.GetModelReply.value)
  return _internal_value();
}
template <typename Arg_, typename... Args_>
inline PROTOBUF_ALWAYS_INLINE void GetModelReply::set_value(Arg_&& arg,
                                                     Args_... args) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.value_.SetBytes(static_cast<Arg_&&>(arg), args..., GetArena());
  // @@protoc_insertion_point(field_set:helloworld.v1.GetModelReply.value)
}
inline std::string* GetModelReply::mutable_value() ABSL_ATTRIBUTE_LIFETIME_BOUND {
  std::string* _s = _internal_mutable_value();
  // @@protoc_insertion_point(field_mutable:helloworld.v1.GetModelReply.value)
  return _s;
}
inline const std::string& GetModelReply::_internal_value() const {
  ::google::protobuf::internal::TSanRead(&_impl_);
  return _impl_.value_.Get();
}
inline void GetModelReply::_internal_set_value(const std::string& value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.value_.Set(value, GetArena());
}
inline std::string* GetModelReply::_internal_mutable_value() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  return _impl_.value_.Mutable( GetArena());
}
inline std::string* GetModelReply::release_value() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  // @@protoc_insertion_point(field_release:helloworld.v1.GetModelReply.value)
  return _impl_.value_.Release();
}
inline void GetModelReply::set_allocated_value(std::string* value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.value_.SetAllocated(value, GetArena());
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        if (_impl_.value_.IsDefault()) {
          _impl_.value_.Set("", GetArena());
        }
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:helloworld.v1.GetModelReply.value)
}

#ifdef __GNUC__
#pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace helloworld


// @@protoc_insertion_point(global_scope)

#include "google/protobuf/port_undef.inc"

#endif  // GOOGLE_PROTOBUF_INCLUDED_huajian_2fv1_2fhuajian_2eproto_2epb_2eh